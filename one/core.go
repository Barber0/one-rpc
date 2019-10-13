package one

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"one/transport"
	"one/util/logger"
	"sync"
	"time"
)

var (
	globalLogger	=	logger.GetOneLogger("global")
	rpcSvr			=	make(map[string]*transport.OneSvr)
	svrWg			sync.WaitGroup

	ConfPath		string
	globalConf		struct {
		Server		map[string]*transport.OneSvrConf	`yaml:"server"`
	}
)

func init() {
	flag.StringVar(&ConfPath,"config","","config path")
	Init()
}

func Init() {
	cfgFile,_ := ioutil.ReadFile(ConfPath)
	yaml.Unmarshal(cfgFile, &globalConf)
	for name, svr := range globalConf.Server {
		if svr.AcceptTimeout != 0 {
			globalConf.Server[name].AcceptTimeout = svr.AcceptTimeout * time.Millisecond
		}else {
			globalConf.Server[name].AcceptTimeout = AcceptTimeout
		}
		globalConf.Server[name].ReadTimeout = svr.ReadTimeout * time.Millisecond
		globalConf.Server[name].WriteTimeout = svr.WriteTimeout * time.Millisecond
		globalConf.Server[name].HandleTimeout = svr.HandleTimeout * time.Millisecond
		globalConf.Server[name].IdleTimeout = svr.IdleTimeout * time.Millisecond

		if svr.QueueCap == 0 {
			globalConf.Server[name].QueueCap = QueueCap
		}
		if svr.TCPReadBuf == 0{
			globalConf.Server[name].TCPReadBuf = TCPReadBuf
		}
		if svr.TCPWriteBuf == 0{
			globalConf.Server[name].TCPWriteBuf = TCPWriteBuf
		}
	}
}

func Run() {
	for obj, svr := range rpcSvr {
		svrWg.Add(1)
		go func() {
			defer svrWg.Done()
			if err := svr.Serve(); err != nil {
				globalLogger.Errorf("server %s err: %v",obj,err)
			}
		}()
	}
	svrWg.Wait()
}