package one

import (
	"flag"
	"fmt"
	"github.com/Barber0/one-rpc/registry"
	"github.com/Barber0/one-rpc/transport"
	"github.com/Barber0/one-rpc/util/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
	"time"
)

var (
	ctx			Context
	ConfPath	string
)

type Context *contextImp

type contextImp struct {
	Logger		logger.Logger
	RpcSvr		map[string]*transport.OneSvr
	svrWg		sync.WaitGroup
	Conf		OneGlobalConf
}

type OneGlobalConf struct {
	LogPath		string								`yaml:"logpath"`
	Server		map[string]*transport.OneSvrConf	`yaml:"server"`
	Client		transport.OneCltConf				`yaml:"client"`
	Registry	*registry.RegistryConf				`yaml:"registry"`
}

func init() {
	flag.StringVar(&ConfPath,"config","config.yaml","config path")
	ctx = &contextImp{
		Logger:		logger.GetOneLogger("global"),
		RpcSvr:		make(map[string]*transport.OneSvr),
	}
	Init()
}

func Init() {
	cfgFile,err := ioutil.ReadFile(ConfPath)
	if err != nil {
		panic(fmt.Errorf("parse config err: %v, path: %v", err, ConfPath))
	}
	if err = yaml.Unmarshal(cfgFile, &ctx.Conf); err != nil {
		panic(fmt.Errorf("parse config err: %v", err))
	}
	for name, svr := range ctx.Conf.Server {
		if svr.AcceptTimeout != 0 {
			ctx.Conf.Server[name].AcceptTimeout = svr.AcceptTimeout * time.Millisecond
		}else {
			ctx.Conf.Server[name].AcceptTimeout = SvrAcceptTimeout
		}
		ctx.Conf.Server[name].ReadTimeout = svr.ReadTimeout * time.Millisecond
		ctx.Conf.Server[name].WriteTimeout = svr.WriteTimeout * time.Millisecond
		ctx.Conf.Server[name].HandleTimeout = svr.HandleTimeout * time.Millisecond
		ctx.Conf.Server[name].IdleTimeout = svr.IdleTimeout * time.Millisecond

		if svr.QueueCap == 0 {
			ctx.Conf.Server[name].QueueCap = QueueCap
		}
		if svr.TCPReadBuf == 0{
			ctx.Conf.Server[name].TCPReadBuf = TCPReadBuf
		}
		if svr.TCPWriteBuf == 0{
			ctx.Conf.Server[name].TCPWriteBuf = TCPWriteBuf
		}
	}
	if ctx.Conf.Client.Balance == "" {
		ctx.Conf.Client.Balance = NORMAL_BALANCE
	}
	if ctx.Conf.Client.TransProtocol == "" {
		ctx.Conf.Client.TransProtocol = PROTOCOL
	}
	if ctx.Conf.Client.DialTimeout != 0 {
		ctx.Conf.Client.DialTimeout *= time.Millisecond
	}else {
		ctx.Conf.Client.DialTimeout = CltDialTimeout
	}
	if ctx.Conf.Client.ReadTimeout != 0 {
		ctx.Conf.Client.ReadTimeout *= time.Millisecond
	}else {
		ctx.Conf.Client.ReadTimeout = CltReadTimeout
	}
	if ctx.Conf.Client.WriteTimeout != 0 {
		ctx.Conf.Client.WriteTimeout *= time.Millisecond
	}else {
		ctx.Conf.Client.WriteTimeout = CltWriteTimeout
	}
	if ctx.Conf.Client.IdleTimeout != 0 {
		ctx.Conf.Client.IdleTimeout *= time.Millisecond
	}else {
		ctx.Conf.Client.IdleTimeout = CltIdleTimeout
	}
}

func GetContext() Context {
	return ctx
}

func Run() {
	for obj, svr := range ctx.RpcSvr {
		ctx.svrWg.Add(1)
		go func() {
			defer ctx.svrWg.Done()
			if err := svr.Serve(); err != nil {
				ctx.Logger.Errorf("server %s err: %v",obj,err)
			}
		}()
	}
	ctx.svrWg.Wait()
}