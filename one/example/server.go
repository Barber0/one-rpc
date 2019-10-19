package main

import (
	"context"
	"errors"
	"one"
	"one/example/res/test"
)

type ServiceImp struct {}

func (imp *ServiceImp) Alpha(ctx context.Context, req *test.Obj) (rsp *test.Obj, err error) {
	if req.Code > 100 {
		err = errors.New("ffffffff")
		return
	}
	rsp = &test.Obj{
		Code:		req.Code + 100,
	}
	return
}

func main() {
	app := new(test.AppService)
	imp := new(ServiceImp)
	app.RegisterServiceImp("alpha",imp)
	one.Run()
}
