package one

import (
	"context"
	"one/balance"
	"one/transport"
	"sync/atomic"
)

const MAX_INT int32 = 1<<31 - 1

type ServiceController struct {
	reqId		int32
	alias		string
	balancer	balance.Balancer
}

func NewServiceController(alias string, addrs ...string) *ServiceController {
	sc := &ServiceController{
		alias:		alias,
	}
	switch globalConf.Client.Balance {
	case NORMAL_BALANCE:
		sc.balancer = balance.NewNormalBalancer()
	}

	protos := make([]balance.Node,len(addrs))
	for i,addr := range addrs {
		cfg := globalConf.Client
		cfg.Address = addr
		proto := &OneClientProtocol{}
		proto.clt = transport.NewOneClt(proto,globalLogger,&cfg)
		protos[i] = proto
	}
	sc.balancer.Add(protos...)

	return sc
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