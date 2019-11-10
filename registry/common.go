package registry

import (
	"fmt"
	"net"
)

type RegisterCenter interface {
	Register(name string, meta *AppMeta) error
	UnRegister(name string) error
	GetServices(name string) ([]AppMeta, error)
}

type RegistryConf struct {
	Etcd	*EtcdConf		`yaml:"etcd"`
}

type AppMeta struct {
	IP			string	`json:"ip"`
	Port		int		`json:"port"`
	Weight		int		`json:"weight"`
	MetaData	string	`json:"meta_data"`
}

func GetLocalIP() (ip string, err error) {
	var addrs []net.Addr
	addrs, err = net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		if ipAddr, ok := addr.(*net.IPNet); ok && !ipAddr.IP.IsLoopback() {
			if ipv4 := ipAddr.IP.To4(); ipv4 != nil {
				ip = ipv4.String()
				return
			}
		}
	}
	err = fmt.Errorf("failed to fetch ip")
	return
}