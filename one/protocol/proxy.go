package protocol

import (
	"one/transport"
	"one/util/logger"
)

func AddProxy(name string, dpr dispatcher, imp interface{}, regFunc func(name string, sp transport.SvrProtocol) error) error {
	p := NewServerProtocol(dpr,imp,logger.GetOneLogger(name))
	return regFunc(name,p)
}
