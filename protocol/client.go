package protocol

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/Barber0/one/protocol/res/requestf"
	"github.com/Barber0/one/transport"
	"github.com/Barber0/one/util/logger"
	"sync"
)

type ClientProtocol struct {
	rsp		sync.Map
	clt		*transport.OneClt
	logger	Logger
}

func NewClientProtocol(name string, conf *transport.OneCltConf) transport.CltProtocol {
	cp := &ClientProtocol{}
	cp.clt = transport.NewOneClt(cp,logger.GetOneLogger(name),conf)
	return cp
}

func (ocp *ClientProtocol) Recv(pkg []byte) {
	rsp := new(requestf.RspPacket)
	proto.Unmarshal(pkg,rsp)
	if readC, ok := ocp.rsp.Load(rsp.ReqId); ok {
		readC.(chan *requestf.RspPacket) <- rsp
	}else {
		ocp.logger.Error("no such message")
	}
}

func (ocp *ClientProtocol) Send(reqId int32, servant, funcName string, pkg []byte) (rspPkg []byte, err error) {
	req := &requestf.ReqPacket{
		Version:	ONE_RPC_VERSION,
		ReqId:		reqId,
		Servant:	servant,
		FuncName:	funcName,
		Content:	pkg,
	}
	reqPkg := req.Bytes()
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

func (ocp *ClientProtocol) ParsePkg(pkg []byte) (int, int) {
	return transport.ParsePkg(pkg)
}

func (ocp *ClientProtocol) getReqId() int32 {
	return 0
}

func (ocp *ClientProtocol) String() string {
	cltCfg := ocp.clt.GetConf()
	return fmt.Sprintf("%s(%s)",cltCfg.TransProtocol,cltCfg.Address)
}