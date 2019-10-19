package test

import (
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/Barber0/one"
	"github.com/Barber0/one/protocol"
	"github.com/Barber0/one/protocol/res/requestf"
)

type AppService struct{
	c *one.ServiceController
}

func NewAppService(addr ...string) *AppService {
	as := &AppService{
		c:		one.NewServiceController("alpha",protocol.NewClientProtocol,addr...),
	}
	return as
}

func (a *AppService) Alpha(ctx context.Context, req *Obj) (rsp *Obj, err error) {
	var (
		reqPkg	[]byte
		rspPkg	[]byte
	)
	reqPkg,_ = proto.Marshal(req)
	if rspPkg,err = a.c.Send(ctx,"AppService","Alpha",reqPkg); err != nil {
		return
	}
	rsp = &Obj{}
	proto.Unmarshal(rspPkg,rsp)
	return
}

type _service interface {
	Alpha(ctx context.Context,req *Obj) (*Obj,error)
}

func (s *AppService) RegisterServiceImp(name string, imp _service) error {
	return protocol.AddProxy(name, s, imp, one.SetRpcServer)
}

// Dispatch方法运行在服务端
// Dispatcher接口由protobuf生成的Service结构体实现
func (d *AppService) Dispatch(ctx context.Context, imp interface{}, req *requestf.ReqPacket, rsp *requestf.RspPacket) (err error) {
	switch req.FuncName {
	case "Alpha":
		payload := new(Obj)
		if err = proto.Unmarshal(req.Content, payload); err != nil {
			return
		}
		var out *Obj
		out,err = imp.(_service).Alpha(ctx,payload)
		if err != nil {
			return
		}
		*rsp = requestf.RspPacket{
			Version:	one.ONE_RPC_VERSION,
			ReqId:		req.ReqId,
		}
		rsp.Content,_ = proto.Marshal(out)
	default:
		err = errors.New("func mismatch, no such func")
	}
	return
}