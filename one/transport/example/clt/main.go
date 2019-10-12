package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"one/transport"
	"one/util/logger"
	"time"
)

type BetaProtocol struct {}

func (bp *BetaProtocol) Recv(pkg []byte) {
	fmt.Println(string(pkg[4:]))
}

func (bp *BetaProtocol) ParsePkg(pkg []byte) (int, int) {
	return transport.ParsePkg(pkg)
}

func main() {
	cltLogger := logger.GetOneLogger("client")
	cltConf := transport.OneCltConf{
		Address:		"127.0.0.1:5000",
		TransProtocol:	"tcp",
		DialTimeout:	3 * time.Minute,
		ReadTimeout:	3 * time.Second,
		WriteTimeout:	3 * time.Second,
		IdleTimeout:	3 * time.Minute,
	}
	svr := transport.NewOneClt(&BetaProtocol{},cltLogger,cltConf)

	buf := bytes.NewBuffer(make([]byte,4))
	buf.Write([]byte("client fff"))
	req := buf.Bytes()
	binary.BigEndian.PutUint32(req,uint32(buf.Len()))

	svr.Send(req)
	time.Sleep(10 * time.Second)
}