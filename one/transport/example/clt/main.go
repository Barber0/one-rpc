package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"one/protocol/res/requestf"
	"one/transport"
	"one/util/logger"
	"runtime/debug"
	"time"
)

type BetaProtocol struct {
	clt		*transport.OneClt
}

func (bp *BetaProtocol) Recv(pkg []byte) {
	defer func() {
		if pa := recover(); pa != nil {
			debug.PrintStack()
			fmt.Println(pa)
		}
	}()
	rsp := new(requestf.RspPacket)
	if err := proto.Unmarshal(pkg, rsp); err != nil {
		panic(err)
	}
	fmt.Println(rsp,string(rsp.Content))
}

func (bp *BetaProtocol) Send(pkg []byte) (rspPkg []byte, err error) {
	return
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
	clt := transport.NewOneClt(&BetaProtocol{},cltLogger,cltConf)

	req := requestf.Req2Bytes(&requestf.ReqPacket{
		Version:	1,
		Content:	[]byte("test"),
	})

	clt.Send(req)
	time.Sleep(10 * time.Second)
}