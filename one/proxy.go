package one

import "one/transport"

type Proxy struct {
	svr		*transport.OneSvr
}

func AddProxy(obj string, dpr dispatcher) {
	p := NewOneProtocol(dpr)
	rpcSvr[obj] = transport.NewOneSvr(p,globalLogger,getSvrConf(obj))
}
