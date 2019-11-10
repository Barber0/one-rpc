package main

import (
	"context"
	"fmt"
	"github.com/Barber0/one-rpc/example/res/test"
	"log"
	"time"
)

func main() {
	app, err := test.NewQnA("alpha")
	if err != nil {
		log.Fatal(err)
	}

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