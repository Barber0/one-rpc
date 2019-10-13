package one_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"one"
	"one/protocol/res/endpointf"
	"one/protocol/res/requestf"
	"testing"
)

// 开发者自行编写
type ServiceImp struct {}

func (imp *ServiceImp) Alpha(ctx context.Context, req *endpointf.EndpointF) (rsp *endpointf.EndpointF, err error) {
	return
}
// end 开发者自行编写

// 框架实现
//func RegisterService(obj string)  {
//
//}
// end 框架实现

// protoc生成
type AppService struct{}

type _service interface {
	Alpha(ctx context.Context,req *endpointf.EndpointF) (*endpointf.EndpointF,error)
}

func (s *AppService) RegisterServiceImp(obj string, imp _service) {
	one.AddProxy(obj, s, imp)
}

// Dispatch方法运行在服务端
// Dispatcher接口由protobuf生成的Service结构体实现
func (d *AppService) Dispatch(ctx context.Context, imp interface{}, req *requestf.ReqPacket, rsp *requestf.RspPacket) (err error) {
	switch req.FuncName {
	case "Alpha":
		payload := new(endpointf.EndpointF)
		if err = proto.Unmarshal(req.Content, payload); err != nil {
			return
		}
		var out *endpointf.EndpointF
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

// end protoc 生成

func TestNewOneProtocol(t *testing.T) {
	p := one.NewOneProtocol(&AppService{},&ServiceImp{})
	fmt.Println(p)
}

func TestService(t *testing.T) {
	app := new(AppService)
	imp := new(ServiceImp)
	app.RegisterServiceImp("alpha",imp)
	one.Run()
}