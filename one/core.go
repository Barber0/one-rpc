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
	globalConf		OneGlobalConf
)

type OneGlobalConf struct {
	Server		map[string]*transport.OneSvrConf	`yaml:"server"`
	Client		transport.OneCltConf				`yaml:"client"`
}

func init() {
	flag.StringVar(&ConfPath,"config","config.yaml","config path")
	Init()
}

func Init() {
	cfgFile,_ := ioutil.ReadFile(ConfPath)
	yaml.Unmarshal(cfgFile, &globalConf)
	for name, svr := range globalConf.Server {
		if svr.AcceptTimeout != 0 {
			globalConf.Server[name].AcceptTimeout = svr.AcceptTimeout * time.Millisecond
		}else {
			globalConf.Server[name].AcceptTimeout = SvrAcceptTimeout
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
	if globalConf.Client.Balance == "" {
		globalConf.Client.Balance = NORMAL_BALANCE
	}
	if globalConf.Client.DialTimeout != 0 {
		globalConf.Client.DialTimeout *= time.Millisecond
	}else {
		globalConf.Client.DialTimeout = CltDialTimeout
	}
	if globalConf.Client.ReadTimeout != 0 {
		globalConf.Client.ReadTimeout *= time.Millisecond
	}else {
		globalConf.Client.ReadTimeout = CltReadTimeout
	}
	if globalConf.Client.WriteTimeout != 0 {
		globalConf.Client.WriteTimeout *= time.Millisecond
	}else {
		globalConf.Client.WriteTimeout = CltWriteTimeout
	}
	if globalConf.Client.IdleTimeout != 0 {
		globalConf.Client.IdleTimeout *= time.Millisecond
	}else {
		globalConf.Client.IdleTimeout = CltIdleTimeout
	}
}

func GetConf() *OneGlobalConf {
	return &globalConf
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