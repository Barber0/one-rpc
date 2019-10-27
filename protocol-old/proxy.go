package protocol

import (
	"errors"
	"github.com/Barber0/one-rpc"
	"github.com/Barber0/one-rpc/transport"
	"github.com/Barber0/one-rpc/util/logger"
)

func AddProxy(name string, dpr dispatcher, imp interface{}) error {
	p := NewServerProtocol(dpr,imp,logger.GetOneLogger(name))
	ctx := one.GetContext()
	if conf, ok := ctx.Conf.Server[name]; !ok {
		return errors.New("fetch server config failed")
	}else {
		ctx.RpcSvr[name] = transport.NewOneSvr(p,logger.GetOneLogger(name),conf)
	}
	return nil
}
