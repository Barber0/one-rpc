package main

import (
	"context"
	"fmt"
	"github.com/Barber0/one-rpc/example/res/test"
	"time"
)

func main() {
	app := test.NewAppService("127.0.0.1:5000")

	tenTimes(app)
	time.Sleep(5*time.Second)
	tenTimes(app)
	tenTimes(app)
}

func tenTimes(app *test.AppService) {
	times := 10
	start := time.Now()
	for i := 0; i < times; i++ {
		fmt.Println(app.Alpha(context.Background(),&test.Obj{
			Code:	int32(i * 10),
		}))
	}
	fmt.Println(time.Now().Sub(start))
}