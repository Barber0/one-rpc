package registry_test

import (
	"context"
	"fmt"
	"github.com/Barber0/one-rpc/registry"
	"testing"
	"time"
)

func TestNewEtcdRegistryCenter(t *testing.T) {
	conf := &registry.EtcdConf{
		[]string{"127.0.0.2:5103"},
		20,
		10,
		9,
	}
	reg, err := registry.NewEtcdRegistryCenter(context.Background(), conf)
	if err != nil {
		t.Error(err)
	}
	ip, err := registry.GetLocalIP()
	if err != nil {
		t.Error(err)
	}
	svrName := "/service/alpha"
	svrMeta := &registry.AppMeta{
		IP:		ip,
		Port:	5000,
	}
	if err = reg.Register(svrName, svrMeta); err != nil {
		t.Error(err)
	}

	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println(reg.UnRegister(svrName))


		time.Sleep(time.Minute)
	}()
}
