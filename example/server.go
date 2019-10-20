package main

import (
	"context"
	"fmt"
	"github.com/Barber0/one-rpc"
	"github.com/Barber0/one-rpc/example/res/test"
)

type ServiceImp struct {}

func (imp *ServiceImp) Ask(ctx context.Context, req *test.Question) (rsp *test.Answer, err error) {
	fmt.Println(req.Msg)
	rsp = &test.Answer{
		Msg:	req.Msg+"----------ddd",
	}
	return
}

func main() {
	app := new(test.QnA)
	imp := new(ServiceImp)
	app.RegisterServiceImp("alpha",imp)
	one.Run()
}
