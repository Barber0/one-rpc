package one

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"one/protocol/res/requestf"
	"one/transport"
	"sync"
)

type OneClientProtocol struct {
	rsp		sync.Map
	clt		*transport.OneClt
}

func (ocp *OneClientProtocol) Recv(pkg []byte) {
	rsp := new(requestf.RspPacket)
	proto.Unmarshal(pkg,rsp)
	if readC, ok := ocp.rsp.Load(rsp.ReqId); ok {
		readC.(chan *requestf.RspPacket) <- rsp
	}else {
		globalLogger.Error("no such message")
	}
}

func (ocp *OneClientProtocol) Send(reqId int32, servant, funcName string, pkg []byte) (rspPkg []byte, err error) {
	req := &requestf.ReqPacket{
		Version:	ONE_RPC_VERSION,
		ReqId:		reqId,
		Servant:	servant,
		FuncName:	funcName,
		Content:	pkg,
	}
	reqPkg := requestf.Req2Bytes(req)
	readC := make(chan *requestf.RspPacket)
	defer func() {
		ocp.rsp.Delete(reqId)
		close(readC)
	}()
	ocp.rsp.Store(reqId, readC)
	if err = ocp.clt.Send(reqPkg); err != nil {
		return
	}
	rsp := <-readC
	if rsp.IsErr {
		err = errors.New(rsp.ResDesc)
	}else {
		rspPkg = rsp.Content
	}
	return
}

func (ocp *OneClientProtocol) ParsePkg(pkg []byte) (int, int) {
	return transport.ParsePkg(pkg)
}

func (ocp *OneClientProtocol) getReqId() int32 {
	return 0
}

func (ocp *OneClientProtocol) String() string {
	cltCfg := ocp.clt.GetConf()
	return fmt.Sprintf("%s(%s)",cltCfg.TransProtocol,cltCfg.Address)
}