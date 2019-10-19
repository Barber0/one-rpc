package main

import (
	"context"
	"fmt"
	"github.com/Barber0/one/example/res/test"
)

func main() {
	app := test.NewAppService("127.0.0.1:5000")
	fmt.Println(app.Alpha(context.Background(),&test.Obj{
		Code:	10,
	}))
}