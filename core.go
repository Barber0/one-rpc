package one

import (
	"context"
	"flag"
	"fmt"
	"github.com/Barber0/one-rpc/registry"
	"github.com/Barber0/one-rpc/transport"
	"github.com/Barber0/one-rpc/util/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	ctx			Context
	ConfPath	string
)

type (
	Context 			*contextImp
)

type contextImp struct {
	Logger		logger.Logger
	RpcSvr		map[string]*transport.OneSvr
	svrWg		sync.WaitGroup
	Conf		OneGlobalConf
	RegisterCenters	map[string]registry.RegisterCenter
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
		RegisterCenters:	make(map[string]registry.RegisterCenter),
	}
	Init()
}

func Init() {
	var err error
	cfgFile,err := ioutil.ReadFile(ConfPath)
	if err != nil {
		panic(fmt.Errorf("parse config err: %v, path: %v", err, ConfPath))
	}
	if err = yaml.Unmarshal(cfgFile, &ctx.Conf); err != nil {
		panic(fmt.Errorf("parse config err: %v", err))
	}
	if etcdConf := ctx.Conf.Registry.Etcd; etcdConf != nil {
		ctx.RegisterCenters[REGISTER_CENTER_ETCD], err = registry.NewEtcdRegistryCenter(context.Background(), etcdConf)
		if err != nil {
			panic(fmt.Errorf("init etcd registry center failed, err: %v", err))
		}
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

func RegisterService(name string, serverConf *transport.OneSvrConf, metaInfo...string) error {
	ip, err := registry.GetLocalIP()
	if err != nil {
		return err
	}
	addr, err := net.ResolveTCPAddr("tcp", serverConf.Address)
	if err != nil {
		return err
	}
	meta := &registry.AppMeta{
		IP:		ip,
		Port:	addr.Port,
		Weight:	serverConf.ServiceWeight,
		MetaData:	strings.Join(metaInfo,"\r\n"),
	}
	for regName, center := range ctx.RegisterCenters {
		if err := center.Register(name, meta); err != nil {
			return err
		}
		ctx.Logger.Debugf("Register %s in %s, meta: %v", name, regName, meta)
	}
	return nil
}