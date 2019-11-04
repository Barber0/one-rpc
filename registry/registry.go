package registry

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Barber0/one-rpc/util/logger"
	cliv3 "go.etcd.io/etcd/clientv3"
	"net"
	"strconv"
	"time"
)

type RegisterService struct {
	Cli			*cliv3.Client
	LeaseRsp	*cliv3.LeaseGrantResponse
	KeepAliveC	<-chan *cliv3.LeaseKeepAliveResponse
	Exit		bool
	ServiceKey	string
	Service		*ServiceInfo
	ErrC		chan error
	logger		*logger.OneLogger
}

type ServiceInfo struct {
	Endpoint	string
	Weight		int
	Room		string
}

func NewRegisterService(endpoints []string, username, password string) (rs *RegisterService, err error) {
	regService := &RegisterService{
		Exit:		false,
		logger:		logger.GetOneLogger("register"),
	}
	var cli *cliv3.Client
	if cli, err = GetClient(endpoints, 5, 3, username, password); err != nil {
		return
	}
	regService.Cli = cli
	regService.ErrC = make(chan error,1)
	rs = regService
	return
}

func (rs *RegisterService) Register(svrName string, svrInfo *ServiceInfo, leaseTime int64) (err error) {
	var (
		leaseRsp	*cliv3.LeaseGrantResponse
		keepAliveC	<-chan *cliv3.LeaseKeepAliveResponse
	)
	leaseRsp, err = rs.Cli.Grant(context.TODO(),leaseTime)
	if err != nil {
		rs.ErrC <- err
		return
	}
	rs.LeaseRsp = leaseRsp
	keepAliveC,err = rs.Cli.KeepAlive(context.TODO(), leaseRsp.ID)
	if err != nil {
		rs.ErrC <- err
		return
	}
	rs.KeepAliveC = keepAliveC
	leaseId := leaseRsp.ID
	go rs.listenLeaseRspC()

	rs.ServiceKey = fmt.Sprintf("%s/%s",svrName, strconv.FormatInt(int64(leaseId),10))
	bsVal, err := json.Marshal(svrInfo)
	if err != nil {
		rs.ErrC <- err
		return
	}
	if _, err = rs.Cli.Put(context.TODO(), rs.ServiceKey, string(bsVal)); err != nil {
		return
	}
	rs.Service = svrInfo
	return
}

func (rs *RegisterService) listenLeaseRspC() {
	for {
		select {
		case leaseKeepAlive := <-rs.KeepAliveC:
			if leaseKeepAlive == nil {
				if !rs.Exit {
					rs.logger.Error("service lease keep alive failed")
					rs.ErrC <- errors.New("invalid revoke")
				}
				return
			}
		}
	}
}

func (rs *RegisterService) UnRegister() (err error) {
	rs.Exit = true
	_,err = rs.Cli.Revoke(context.TODO(),rs.LeaseRsp.ID)
	return
}

func GetClient(endpoints []string, dialTimeout, dialKeepAliveTime int, username, password string) (cli *cliv3.Client, err error) {
	cli,err = cliv3.New(cliv3.Config{
		Endpoints:		endpoints,
		DialTimeout:	time.Duration(dialTimeout) * time.Second,
		DialKeepAliveTime:	time.Duration(dialKeepAliveTime) * time.Second,
		//Username:		username,
		//Password:		password,
	})
	if err != nil {
		cli = nil
	}
	return
}

func GetLocalIP() (ip string, err error) {
	addrs,err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _,addr := range addrs {
		if tmpip, ok := addr.(*net.IPNet); ok && !tmpip.IP.IsLoopback() {
			if tmpip.IP.To4() != nil {
				ip = tmpip.IP.To4().String()
				return
			}
		}
	}
	err = errors.New("no ip address")
	return
}