package registry_test

import (
	"fmt"
	"github.com/Barber0/one-rpc/registry"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestNewRegisterService(t *testing.T) {
	ip,err := registry.GetLocalIP()
	if err != nil {
		t.Error(err)
	}

	svrName := "fang:0"
	svrInfo	:= &registry.ServiceInfo{
		Endpoint:	ip,
		Weight:		100,
	}
	var leaseTime int64	=	10

	endpoints := []string{"127.0.0.1:2380"}
	reg,err := registry.NewRegisterService(endpoints,"alpha","beta")
	if err != nil {
		t.Error(err)
	}
	if err := reg.Register(svrName, svrInfo, leaseTime); err != nil {
		t.Error(err)
	}
	fmt.Println("fffff")
	fmt.Println(reg.ServiceKey)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	//创建一个新的goroutine监听退出信号，在ETCD下线服务并优雅退出服务

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer func() {
			wg.Done()
		}()
		sig := <-stopChan
		fmt.Printf("sig: %v",sig)
		err := reg.UnRegister()
		if err != nil {
			fmt.Println("sevice offline failed:", err)
			t.Error(err)
		}
		//服务处理完当前遗留请求之后，优雅退出
		//server.GracefulStop()
		//os.Exit(0)
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		count := 0
		for {
			select {
			case keepError := <-reg.ErrC:
				count++
				//如果服务续租或者注册失败，进行重新注册保活。
				if keepError != nil {
					err := reg.Register(svrName, svrInfo, leaseTime)
					if err != nil {
						log.Println("service keep alive fail again")
						time.Sleep(3 * time.Second)
					}
				}
			}
			if count > 10 {
				return
			}
			//fmt.Println("count: ",count)
		}

	}()

	fmt.Println("ddd")
	wg.Wait()
	fmt.Println("ddd---fffff")
}
