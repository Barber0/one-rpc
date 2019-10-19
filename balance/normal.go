package balance

import (
	"errors"
	"sync"
	"sync/atomic"
)

var _ Balancer = &NormalBalancer{}

type NormalBalancer struct {
	nodes		[]Node
	position	int32
	hasAdded	int32
	lock		*sync.Mutex
}

func NewNormalBalancer() *NormalBalancer {
	return &NormalBalancer{lock:&sync.Mutex{}}
}

func (nb *NormalBalancer) Add(nodes ...Node) (err error) {
	if atomic.CompareAndSwapInt32(&nb.hasAdded,0,1) {
		nb.nodes = nodes
	}else {
		nb.nodes = append(nb.nodes,nodes...)
		err = errors.New("normal balancer could only be added for one time")
	}
	return
}

func (nb *NormalBalancer) Delete(nodes ...Node) (err error) {
	return errors.New("normal balancer could not be deleted")
}

func (nb *NormalBalancer) GetNode(bs []byte) (res Node, ok bool) {
	nb.lock.Lock()
	defer nb.lock.Unlock()
	nb.position = (nb.position + 1) % int32(len(nb.nodes))
	res = nb.nodes[nb.position]
	ok = true
	return
}