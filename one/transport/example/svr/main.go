package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"one/transport"
	"one/util/logger"
	"time"
)

type AlphaProtocol struct {}

func (ap *AlphaProtocol) Invoke(ctx context.Context, req []byte) (rsp []byte) {
	buf := bytes.NewBuffer(make([]byte,4))
	buf.Write([]byte("server ddd"))
	rsp = buf.Bytes()
	binary.BigEndian.PutUint32(rsp,uint32(buf.Len()))
	return
}

func (ap *AlphaProtocol) InvokeTimeout(ctx context.Context, req []byte) (rsp []byte) {
	return
}

func (ap *AlphaProtocol) ParsePkg(pkg []byte) (int, int) {
	return transport.ParsePkg(pkg)
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
		TCPReadBuf:		4 * 1024 * 1024,
		TCPWriteBuf:	4 * 1024 * 1024,
		TCPNoDelay:		false,
	}
	svr := transport.NewOneSvr(&AlphaProtocol{},svrLogger,svrConf)
	if err := svr.Serve(); err != nil {
		panic(err)
	}
}