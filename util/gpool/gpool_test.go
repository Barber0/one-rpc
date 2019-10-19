package gpool_test

import (
	"fmt"
	"one/util/gpool"
	"testing"
	"time"
)

func TestNewGPool(t *testing.T) {
	task := 10
	p := gpool.NewGPool(5,20)
	go func() {
		for v := 0; v < task; v++ {
			i := v
			p.AddTask(func() {
				time.Sleep(time.Duration(i+2) * time.Second)
				fmt.Println(i)
			})
		}
	}()
	time.Sleep(time.Second)
	p.Release()
	fmt.Println("fff")
	time.Sleep(30*time.Second)
}
