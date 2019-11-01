package main

import (
	"context"
	"github.com/Barber0/one-rpc"
	"github.com/Barber0/one-rpc/example/res/test"
)

type ServiceImp struct {}

func (imp *ServiceImp) Ask(ctx context.Context, req *test.Question) (rsp *test.Answer, err error) {
	rsp = &test.Answer{
		Msg:	req.Msg,
	}
	return
}

func main() {
	app := &test.QnA{}
	imp := new(ServiceImp)
	app.RegisterServiceImp("alpha",imp)
	one.Run()
}
