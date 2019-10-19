package main

import (
	"context"
	"one"
	"one/protocol/res/requestf"
	"one/transport"
	"one/util/logger"
	"time"
)

type Dispatcher struct{}

func (d *Dispatcher) Dispatch(ctx context.Context, req *requestf.ReqPacket, rsp *requestf.RspPacket) error {
	*rsp = requestf.RspPacket{
		Version:	req.Version,
		Content:	[]byte(string(req.Content)+" answer"),
	}
	return nil
}

func main() {
	svrLogger := logger.GetOneLogger("server")
	svrConf := transport.OneSvrConf{
		Address:		"127.0.0.1:5000",
		TransProtocol:	"tcp",
		AcceptTimeout:	3 * time.Minute,
		ReadTimeout:	3 * time.Second,
		WriteTimeout:	3 * time.Second,
		HandleTimeout:	3 * time.Second,
		IdleTimeout:	3 * time.Minute,
		TCPReadBuf:		4 * 1024 * 1024,
		TCPWriteBuf:	4 * 1024 * 1024,
		TCPNoDelay:		false,
	}
	svr := transport.NewOneSvr(one.NewOneProtocol(&Dispatcher{}),svrLogger,svrConf)
	if err := svr.Serve(); err != nil {
		panic(err)
	}
}