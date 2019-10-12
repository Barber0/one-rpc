package transport

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type OneCltConf struct {
	Address			string
	TransProtocol	string

	QueueCap		int

	DialTimeout		time.Duration
	ReadTimeout		time.Duration
	WriteTimeout	time.Duration
	IdleTimeout		time.Duration
}

type cltConn struct {
	clt			*OneClt
	conn		net.Conn
	isClosed	bool
	connLock	*sync.Mutex
	idleTime	time.Time
	numInvoke	int32
}

type OneClt struct {
	conf		OneCltConf
	proto		CltProtocol
	conn		*cltConn
	logger		Logger
	sendQ		chan []byte
}

func NewOneClt(cp CltProtocol, logger Logger, conf OneCltConf) *OneClt {
	clt := &OneClt{
		conf:	conf,
		proto:	cp,
		logger:	logger,
		sendQ:	make(chan []byte,conf.QueueCap),
	}
	clt.conn = &cltConn{
		clt:		clt,
		isClosed:	true,
		connLock:	&sync.Mutex{},
		idleTime:	time.Now(),
	}
	return clt
}

func (c *OneClt) Send(pkg []byte) (err error) {
	if err = c.conn.reConnect(); err != nil {
		c.conn.close()
		return
	}else {
		c.sendQ <- pkg
	}
	return
}

func (c *cltConn) reConnect() (err error) {
	c.connLock.Lock()
	if c.isClosed {
		cfg := c.clt.conf
		if cfg.DialTimeout != 0 {
			c.conn,err = net.DialTimeout(cfg.TransProtocol,cfg.Address,cfg.DialTimeout)
		}else {
			c.conn,err = net.Dial(cfg.TransProtocol,cfg.Address)
		}
		if err != nil {
			c.clt.logger.Errorf("Dial addr %s err: %v",cfg.Address,err)
			c.connLock.Unlock()
			return
		}
		c.isClosed = false
		c.refreshIdleTime()
		go c.send()
		go c.recv()
	}
	c.connLock.Unlock()
	return
}

func (c *cltConn) send() {
	var (
		pkg []byte
		cfg = c.clt.conf
		t 	= time.NewTicker(time.Second)
	)
	for {
		select {
		case pkg = <-c.clt.sendQ:
		case <-t.C:
			if c.numInvoke == 0 && c.idleTime.Before(time.Now()) {
				c.close()
				return
			}
			continue
		}
		c.refreshIdleTime()
		atomic.AddInt32(&c.numInvoke,1)
		if cfg.WriteTimeout != 0 {
			c.conn.SetWriteDeadline(time.Now().Add(cfg.WriteTimeout))
		}
		if _, err := c.conn.Write(pkg); err != nil {
			c.clt.logger.Errorf("send pkg to %v failed: %v",c.conn.RemoteAddr(),err)
			c.close()
			return
		}
	}
}

func (c *cltConn) recv() {
	var (
		cfg		=	c.clt.conf
		buf		=	make([]byte,4*1024)
		curBuf	[]byte
	)
}

func (c *cltConn) close() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.isClosed = true
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *cltConn) refreshIdleTime() {
	c.idleTime = time.Now().Add(c.clt.conf.IdleTimeout)
}