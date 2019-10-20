package main

import (
	"context"
	"fmt"
	"github.com/Barber0/one-rpc/example/res/test"
	"time"
)

func main() {
	app := test.NewQnA("127.0.0.1:5000")

	tenTimes(app)
	time.Sleep(5*time.Second)
	tenTimes(app)
	tenTimes(app)
}

func tenTimes(app *test.QnA) {
	times := 10
	start := time.Now()
	for i := 0; i < times; i++ {
		fmt.Println(app.Ask(context.Background(),&test.Question{
			Msg:	"hahaha",
		}))
	}
	fmt.Println(time.Now().Sub(start))
}