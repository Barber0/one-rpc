package one

import (
	"one/transport"
	"one/util/logger"
	"sync"
)

var (
	globalLogger	=	logger.GetOneLogger("global")
	rpcSvr			=	make(map[string]*transport.OneSvr)
	svrWg			sync.WaitGroup
)

func getSvrConf(obj string) *transport.OneSvrConf {
	return nil
}

func Run() {
	for obj, svr := range rpcSvr {
		svrWg.Add(1)
		go func() {
			if err := svr.Serve(); err != nil {
				globalLogger.Errorf("server %s err: %v",obj,err)
			}
			svrWg.Done()
		}()
	}
	svrWg.Wait()
}