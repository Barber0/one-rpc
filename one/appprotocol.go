package one

import (
	"context"
	"github.com/golang/protobuf/proto"
	"one/protocol/res/requestf"
	"one/transport"
)

type dispatcher interface {
	Dispatch(ctx context.Context, req *requestf.ReqPacket, rsp *requestf.RspPacket) error
}

type OneProtocol struct {
	dpr		dispatcher
}

func NewOneProtocol(dpr dispatcher) *OneProtocol {
	return &OneProtocol{
		dpr:	dpr,
	}
}

func (p *OneProtocol) Invoke(ctx context.Context, req []byte) []byte {
	reqPacket := new(requestf.ReqPacket)
	rspPacket := new(requestf.RspPacket)
	proto.Unmarshal(req, reqPacket)
	if err := p.dpr.Dispatch(ctx, reqPacket, rspPacket); err != nil {
		globalLogger.Errorf("dispatch request failed: %v",err)
		rspPacket.Version = reqPacket.Version
		rspPacket.ReqId = reqPacket.ReqId
		rspPacket.IsErr	= true
		rspPacket.ResDesc = err.Error()
	}
	return requestf.Rsp2Bytes(rspPacket)
}

func (p *OneProtocol) ParsePkg(pkg []byte) (int, int) {
	return transport.ParsePkg(pkg)
}

func (p *OneProtocol) InvokeTimeout(ctx context.Context, req []byte) []byte {
	reqPacket := new(requestf.ReqPacket)
	proto.Unmarshal(req, reqPacket)
	rspPacket := &requestf.RspPacket{
		Version:	reqPacket.Version,
		ReqId:		reqPacket.ReqId,
		IsErr:		true,
		ResDesc:	"invoke timeout",
	}
	return requestf.Rsp2Bytes(rspPacket)
}