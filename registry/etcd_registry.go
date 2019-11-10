package registry

import (
	"context"
	"encoding/json"
	"fmt"
	cliv3 "go.etcd.io/etcd/clientv3"
	"time"
)

type EtcdConf struct {
	Endpoints	[]string		`yaml:"endpoints"`
	DialTimeout	time.Duration	`yaml:"dial_timeout"`
	DialKeepAliveTime	time.Duration	`yaml:"dial_keepalive_time"`
	LeaseTTL	int64			`yaml:"lease_ttl"`
}

type EtcdRegistryCenter struct {
	ctx			context.Context
	cli			*cliv3.Client
	conf		*EtcdConf
	leases		map[string]*leaseMeta
	ErrC		chan error
}

type leaseMeta struct {
	keepAliveC	<-chan *cliv3.LeaseKeepAliveResponse
	id			cliv3.LeaseID
	appMeta		*AppMeta
}

func NewEtcdRegistryCenter(ctx context.Context, conf *EtcdConf) (reg *EtcdRegistryCenter, err error) {
	reg = &EtcdRegistryCenter{
		ctx:	ctx,
		conf:	conf,
		leases:	make(map[string]*leaseMeta),
		ErrC:	make(chan error),
	}
	reg.cli, err = cliv3.New(cliv3.Config{
		Endpoints:		reg.conf.Endpoints,
		DialTimeout:	reg.conf.DialTimeout * time.Second,
		DialKeepAliveTime:	reg.conf.DialKeepAliveTime * time.Second,
	})
	return
}

func (reg *EtcdRegistryCenter) Register(name string, meta *AppMeta) error {
	leaseRsp, err := reg.cli.Grant(reg.ctx, reg.conf.LeaseTTL)
	if err != nil {
		return err
	}
	reg.leases[name] = &leaseMeta{id: leaseRsp.ID, appMeta:	meta}
	meta_json, _ := json.Marshal(meta)
	if _, err = reg.cli.Put(reg.ctx, name, string(meta_json), cliv3.WithLease(leaseRsp.ID)); err != nil {
		return err
	}
	return reg.keepAlive(name)
}

func (reg *EtcdRegistryCenter) UnRegister(name string) (err error) {
	if meta, ok := reg.leases[name]; ok {
		if _, err = reg.cli.Revoke(reg.ctx, meta.id); err != nil {
			return
		}
		delete(reg.leases, name)
		return
	}
	return fmt.Errorf("failed to get meta in local, service: %s", name)
}

func (reg *EtcdRegistryCenter) GetServices(name string) (metas []AppMeta, err error) {
	var getRsp	*cliv3.GetResponse
	if getRsp, err = reg.cli.Get(reg.ctx, name); err != nil {
		return
	}
	metas = make([]AppMeta, len(getRsp.Kvs))
	tmpMeta := new(AppMeta)
	for i, kv := range getRsp.Kvs {
		if err = json.Unmarshal(kv.Value, tmpMeta); err != nil {
			return
		}
		metas[i] = *tmpMeta
	}
	return
}

func (reg *EtcdRegistryCenter) keepAlive(name string) (err error) {
	if meta, ok := reg.leases[name]; ok {
		if reg.leases[name].keepAliveC, err = reg.cli.KeepAlive(reg.ctx, meta.id); err != nil {
			return
		}
		go func() {
			for ka := range reg.leases[name].keepAliveC {
				if ka == nil {
					reg.ErrC <- fmt.Errorf("keepalive failed, service: %s, lease: %d",name, meta.id)
				}
			}
		}()
		return
	}
	return fmt.Errorf("failed to get meta in local, service: %s", name)
}
