package protocol

import (
	"context"
	"github.com/golang/protobuf/proto"
	"one/protocol/res/requestf"
	"one/transport"
)

type dispatcher interface {
	Dispatch(ctx context.Context, imp interface{}, req *requestf.ReqPacket, rsp *requestf.RspPacket) error
}

type ServerProtocl struct {
	serviceImp	interface{}
	dpr			dispatcher
	logger		Logger
}

func NewServerProtocol(dpr dispatcher, imp interface{}, logger Logger) *ServerProtocl {
	return &ServerProtocl{
		dpr:		dpr,
		serviceImp:	imp,
		logger:		logger,
	}
}

func (p *ServerProtocl) Invoke(ctx context.Context, req []byte) []byte {
	reqPacket := new(requestf.ReqPacket)
	rspPacket := new(requestf.RspPacket)
	proto.Unmarshal(req, reqPacket)
	if err := p.dpr.Dispatch(ctx, p.serviceImp, reqPacket, rspPacket); err != nil {
		p.logger.Errorf("dispatch request failed: %v",err)
		rspPacket.Version = ONE_RPC_VERSION
		rspPacket.ReqId = reqPacket.ReqId
		rspPacket.IsErr	= true
		rspPacket.ResDesc = err.Error()
	}
	return rspPacket.Bytes()
}

func (p *ServerProtocl) ParsePkg(pkg []byte) (int, int) {
	return transport.ParsePkg(pkg)
}

func (p *ServerProtocl) InvokeTimeout(ctx context.Context, req []byte) []byte {
	reqPacket := new(requestf.ReqPacket)
	proto.Unmarshal(req, reqPacket)
	rspPacket := &requestf.RspPacket{
		Version:	ONE_RPC_VERSION,
		ReqId:		reqPacket.ReqId,
		IsErr:		true,
		ResDesc:	"invoke timeout",
	}
	return rspPacket.Bytes()
}
