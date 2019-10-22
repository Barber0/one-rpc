package balance

import (
	"fmt"
	"sort"
	"sync"
)

var _ Balancer = &ConsistentHash{}

type ConsistentHash struct {
	replicate	int
	ring		[]uint32
	tmpRing		[]uint32
	nodes		map[uint32]Node
	valMark		map[Node]struct{}
	hashFunc	func([]byte) uint32
	lock		*sync.Mutex
}

func NewConsistentHash(rep int, hashFunc func([]byte) uint32) *ConsistentHash {
	ch := &ConsistentHash{
		replicate:		rep,
		hashFunc:		hashFunc,
		tmpRing:		make([]uint32,rep),
		nodes:			make(map[uint32]Node),
		valMark:		make(map[Node]struct{}),
		lock:			&sync.Mutex{},
	}
	return ch
}

func (ch *ConsistentHash) Add(nodes ...Node) (err error) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	for _,n := range nodes {
		if _, ok := ch.valMark[n]; ok {
			continue
		}
		for i := 0; i < ch.replicate; i++ {
			hashKey := ch.hashFunc([]byte(fmt.Sprintf("%s#%d",n.String(),i)))
			ch.tmpRing[i] = hashKey
			ch.nodes[hashKey] = n
		}
		ch.ring = append(ch.ring,ch.tmpRing...)
		ch.valMark[n] = struct{}{}
	}
	ch.sortRing()
	return
}

func (ch *ConsistentHash) Delete(nodes ...Node) (err error) {
	ch.lock.Lock()
	defer ch.lock.Unlock()
	for _,n := range nodes {
		for i := 0; i < ch.replicate; i++ {
			hashKey := ch.hashFunc([]byte(fmt.Sprintf("%s#%d",n.String(),i)))
			idx := sort.Search(len(ch.ring), func(j int) bool {
				return ch.ring[j] >= hashKey
			})
			ch.ring = append(ch.ring[:idx],ch.ring[idx+1:]...)
			delete(ch.nodes,hashKey)
		}
		delete(ch.valMark,n)
	}
	return
}

func (ch *ConsistentHash) GetNode(v []byte) (res Node, ok bool) {
	ch.lock.Lock()
	defer ch.lock.Unlock()
	targetVal := ch.hashFunc(v)
	length := len(ch.ring)
	idx := sort.Search(length, func(i int) bool {
		return ch.ring[i] >= targetVal
	})
	if idx < length {
		res, ok = ch.nodes[ch.ring[idx]]
	}else {
		res = ch.nodes[ch.ring[0]]
		ok = true
	}
	return
}

func (ch *ConsistentHash) Show() {
	for _,i := range ch.ring {
		fmt.Println(i,"\t",ch.nodes[i].String())
	}
}

func (ch *ConsistentHash) sortRing() {
	sort.Slice(ch.ring, func(i, j int) bool {
		return ch.ring[i] < ch.ring[j]
	})
}