package transport

import (
	"context"
	"io"
	"net"
	"reflect"
	"time"
)

type tcpHandler struct {
	svr		*OneSvr
	lis		*net.TCPListener
	gpool	GPool
}

func newTcpHandler(svr *OneSvr) *tcpHandler {
	h := &tcpHandler{
		svr:	svr,
	}
	cfg := h.svr.conf
	if cfg.maxInvoke != 0 && cfg.queueCap != 0 {
		h.gpool = NewGPool(cfg.maxInvoke,cfg.queueCap)
	}
	return h
}

func (h *tcpHandler) Listen() (err error) {
	cfg := h.svr.conf
	var addr	*net.TCPAddr
	if addr, err = net.ResolveTCPAddr("tcp4", cfg.Address); err != nil {
		return
	}
	if h.lis, err = net.ListenTCP("tcp4", addr); err != nil {
		return
	}
	h.svr.logger.Debugf("Listening on %s",addr.String())
	return
}

func (h *tcpHandler) Handle() (err error) {
	var conn *net.TCPConn
	cfg := h.svr.conf
	for !h.svr.isClosed {
		h.lis.SetDeadline(time.Now().Add(cfg.AcceptTimeout))
		conn,err = h.lis.AcceptTCP()
		if err != nil {
			if !isNoDataErr(err) {
				h.svr.logger.Errorf("Accept error: %v",err)
			}
			if conn != nil {
				conn.SetKeepAlive(true)
			}
			continue
		}
		go func() {
			h.svr.logger.Debugf("TCP Accept: %v",conn.RemoteAddr())
			conn.SetReadBuffer(cfg.TCPReadBuf)
			conn.SetWriteBuffer(cfg.TCPWriteBuf)
			conn.SetNoDelay(cfg.TCPNoDelay)
			h.recv(conn)
		}()
	}
	if h.gpool != nil {
		h.gpool.Release()
	}
	return
}

func (h *tcpHandler) recv(conn *net.TCPConn) {
	defer conn.Close()
	var (
		cfg		=	h.svr.conf
		buf		=	make([]byte,4*1024)
		curBuf	[]byte
		n		int
		err		error
	)
	for !h.svr.isClosed {
		if cfg.ReadTimeout != 0 {
			conn.SetReadDeadline(time.Now().Add(cfg.ReadTimeout))
		}
		n,err = conn.Read(buf)
		if err != nil {
			if isNoDataErr(err) {
				continue
			}
			if err == io.EOF {
				h.svr.logger.Errorf("conn closed by remote: %v",conn.RemoteAddr())
			}else {
				h.svr.logger.Errorf("read pkg err: %v %v",reflect.TypeOf(err),err)
			}
			return
		}
		curBuf = append(curBuf, buf[:n]...)
		for {
			if pkgLen, status := h.svr.proto.ParsePkg(curBuf); status == PKG_LESS {
				continue
			}else if status == PKG_FULL {
				pkg := make([]byte, pkgLen-4)
				copy(pkg,curBuf[4:pkgLen])
				h.handle(conn, pkg)
				curBuf = curBuf[pkgLen:]
				if len(curBuf) > 0 {
					continue
				}
				curBuf = nil
				break
			}
			h.svr.logger.Errorf("parse pkg err: %v %v",conn.RemoteAddr(),err)
			return
		}
	}
}

func (h *tcpHandler) handle(conn *net.TCPConn, pkg []byte) {
	cfg := h.svr.conf
	handler := func() {
		ctx := context.Background()
		rsp := h.svr.invoke(ctx, pkg)
		if cfg.WriteTimeout != 0 {
			conn.SetWriteDeadline(time.Now().Add(cfg.WriteTimeout))
		}
		if _, err := conn.Write(rsp); err != nil {
			h.svr.logger.Errorf("send pkg to %v err: %v",conn.RemoteAddr(),err)
		}
	}
	if h.gpool != nil {
		h.gpool.AddTask(handler)
	}else {
		go handler()
	}
}