package transport

import (
	"context"
	"sync/atomic"
	"time"
)

type OneSvrConf struct {
	Address			string
	TransProtocol	string

	MaxInvoke		int
	QueueCap		int

	AcceptTimeout	time.Duration
	ReadTimeout		time.Duration
	WriteTimeout	time.Duration
	HandleTimeout	time.Duration
	IdleTimeout		time.Duration

	TCPReadBuf		int
	TCPWriteBuf		int
	TCPNoDelay		bool
}

type OneHandler interface {
	Listen() error
	Handle() error
}

type OneSvr struct {
	isClosed	bool
	handler		OneHandler
	conf		OneSvrConf
	proto		SvrProtocol
	logger		Logger

	numInvoke	int32
	lastInvoke	time.Time
}

func NewOneSvr(sp SvrProtocol, logger Logger, conf OneSvrConf) *OneSvr {
	s := &OneSvr{
		conf:	conf,
		logger:	logger,
		proto:	sp,
	}
	return s
}

func (s *OneSvr) getHandler() (h OneHandler) {
	switch s.conf.TransProtocol {
	case "tcp":
		h = newTcpHandler(s)
	default:
		s.logger.Error("no such protocol")
	}
	return
}

func (s *OneSvr) Serve() (err error) {
	h := s.getHandler()
	if err = h.Listen(); err == nil {
		err = h.Handle()
	}
	return
}

func (s *OneSvr) Shutdown() {
	s.isClosed = true
}

func (s *OneSvr) GetConf() *OneSvrConf {
	return &s.conf
}

func (s *OneSvr) invoke(ctx context.Context, pkg []byte) (rsp []byte) {
	cfg := s.conf
	atomic.AddInt32(&s.numInvoke,1)
	if cfg.HandleTimeout != 0 {
		done := make(chan struct{})
		go func() {
			rsp = s.proto.Invoke(ctx,pkg)
			done <- struct{}{}
		}()
		select {
		case <-done:
		case <-time.After(cfg.HandleTimeout):
			rsp = s.proto.InvokeTimeout(ctx,pkg)
		}
	}else {
		rsp = s.proto.Invoke(ctx,pkg)
	}
	atomic.AddInt32(&s.numInvoke,-1)
	return
}