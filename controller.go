package one

import (
	"context"
	"fmt"
	"github.com/Barber0/one-rpc/balance"
	"github.com/Barber0/one-rpc/registry"
	"github.com/Barber0/one-rpc/transport"
	"sync/atomic"
)

const MAX_INT int32 = 1<<31 - 1

type ClientConstructor func(name string, conf *transport.OneCltConf) transport.CltProtocol

type ServiceController struct {
	reqId		int32
	name		string
	balancer	balance.Balancer
	construct	ClientConstructor
}

var registerCenter	=	REGISTER_CENTER_ETCD

func SetRegisterCenter(regName string) error {
	if _, ok := ctx.RegisterCenters[regName]; ok {
		registerCenter = regName
	}
	return fmt.Errorf("failed to find register center: %s", regName)
}

func NewServiceController(name string, constructor ClientConstructor, addrs ...string) (sc *ServiceController, err error) {
	sc = &ServiceController{
		name:			name,
		construct:		constructor,
	}
	ctx := GetContext()
	conf := ctx.Conf
	switch conf.Client.Balance {
	case NORMAL_BALANCE:
		sc.balancer = balance.NewNormalBalancer()
	}

	if center, ok := ctx.RegisterCenters[registerCenter]; ok {
		var appMetas []registry.AppMeta
		appMetas, err = center.GetServices(name)
		if err != nil {
			return
		}
		ctx.Logger.Debugf("find servers from %s, servers: %v", registerCenter, appMetas)
		for _, meta := range appMetas {
			addrs = append(addrs, fmt.Sprintf("%s:%d",meta.IP,meta.Port))
		}
	}

	protos := make([]balance.Node,len(addrs))
	for i,addr := range addrs {
		cfg := conf.Client
		cfg.Address = addr
		proto := sc.construct(name, &cfg)
		protos[i] = proto
	}
	sc.balancer.Add(protos...)
	return
}

func (sc *ServiceController) Send(ctx context.Context, servant string, funcName string, pkg []byte) (rspPkg []byte, err error) {
	atomic.CompareAndSwapInt32(&sc.reqId,MAX_INT,1)
	msgBs := append([]byte(funcName),pkg...)
	clt := sc.getClient(msgBs)
	return clt.Send(atomic.AddInt32(&sc.reqId,1),servant,funcName,pkg)
}

func (sc *ServiceController) getClient(pkg []byte) (clt transport.CltProtocol) {
	tmpClt,_ := sc.balancer.GetNode(pkg)
	clt = tmpClt.(transport.CltProtocol)
	return
}