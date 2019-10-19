package one

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/Barber0/one/transport"
	"github.com/Barber0/one/util/logger"
	"sync"
	"time"
)

var (
	ctx			*Context
	ConfPath	string
)

type Context struct {
	logger		*logger.OneLogger
	rpcSvr		map[string]*transport.OneSvr
	svrWg		sync.WaitGroup
	conf		OneGlobalConf
}

type OneGlobalConf struct {
	Server		map[string]*transport.OneSvrConf	`yaml:"server"`
	Client		transport.OneCltConf				`yaml:"client"`
}

func init() {
	flag.StringVar(&ConfPath,"config","config.yaml","config path")
	ctx = &Context{
		logger:		logger.GetOneLogger("global"),
		rpcSvr:		make(map[string]*transport.OneSvr),
	}
	Init()
}

func Init() {
	cfgFile,_ := ioutil.ReadFile(ConfPath)
	yaml.Unmarshal(cfgFile, &ctx.conf)
	for name, svr := range ctx.conf.Server {
		if svr.AcceptTimeout != 0 {
			ctx.conf.Server[name].AcceptTimeout = svr.AcceptTimeout * time.Millisecond
		}else {
			ctx.conf.Server[name].AcceptTimeout = SvrAcceptTimeout
		}
		ctx.conf.Server[name].ReadTimeout = svr.ReadTimeout * time.Millisecond
		ctx.conf.Server[name].WriteTimeout = svr.WriteTimeout * time.Millisecond
		ctx.conf.Server[name].HandleTimeout = svr.HandleTimeout * time.Millisecond
		ctx.conf.Server[name].IdleTimeout = svr.IdleTimeout * time.Millisecond

		if svr.QueueCap == 0 {
			ctx.conf.Server[name].QueueCap = QueueCap
		}
		if svr.TCPReadBuf == 0{
			ctx.conf.Server[name].TCPReadBuf = TCPReadBuf
		}
		if svr.TCPWriteBuf == 0{
			ctx.conf.Server[name].TCPWriteBuf = TCPWriteBuf
		}
	}
	if ctx.conf.Client.Balance == "" {
		ctx.conf.Client.Balance = NORMAL_BALANCE
	}
	if ctx.conf.Client.DialTimeout != 0 {
		ctx.conf.Client.DialTimeout *= time.Millisecond
	}else {
		ctx.conf.Client.DialTimeout = CltDialTimeout
	}
	if ctx.conf.Client.ReadTimeout != 0 {
		ctx.conf.Client.ReadTimeout *= time.Millisecond
	}else {
		ctx.conf.Client.ReadTimeout = CltReadTimeout
	}
	if ctx.conf.Client.WriteTimeout != 0 {
		ctx.conf.Client.WriteTimeout *= time.Millisecond
	}else {
		ctx.conf.Client.WriteTimeout = CltWriteTimeout
	}
	if ctx.conf.Client.IdleTimeout != 0 {
		ctx.conf.Client.IdleTimeout *= time.Millisecond
	}else {
		ctx.conf.Client.IdleTimeout = CltIdleTimeout
	}
}

func GetContext() *Context {
	return ctx
}

func SetRpcServer(name string, sp transport.SvrProtocol) error {
	if conf, ok := ctx.conf.Server[name]; !ok {
		return errors.New("fetch server config failed")
	}else {
		ctx.rpcSvr[name] = transport.NewOneSvr(sp,logger.GetOneLogger(name),conf)
	}
	return nil
}

func Run() {
	for obj, svr := range ctx.rpcSvr {
		ctx.svrWg.Add(1)
		go func() {
			defer ctx.svrWg.Done()
			if err := svr.Serve(); err != nil {
				ctx.logger.Errorf("server %s err: %v",obj,err)
			}
		}()
	}
	ctx.svrWg.Wait()
}