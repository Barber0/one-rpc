package protocol

import (
	"github.com/Barber0/one/transport"
	"github.com/Barber0/one/util/logger"
)

func AddProxy(name string, dpr dispatcher, imp interface{}, regFunc func(name string, sp transport.SvrProtocol) error) error {
	p := NewServerProtocol(dpr,imp,logger.GetOneLogger(name))
	return regFunc(name,p)
}
