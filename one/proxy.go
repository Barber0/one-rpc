package one

import (
	"errors"
	"one/transport"
)

type Proxy struct {
	svr		*transport.OneSvr
}

func AddProxy(name string, dpr dispatcher, imp interface{}) {
	var (
		conf 	*transport.OneSvrConf
		ok		bool
	)
	if conf, ok = globalConf.Server[name]; !ok {
		globalLogger.Error(errors.New("fetch server config failed"))
		return
	}
	p := NewOneProtocol(dpr, imp)
	rpcSvr[name] = transport.NewOneSvr(p,globalLogger,conf)
}
