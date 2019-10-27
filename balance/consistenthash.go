package balance

import (
	"fmt"
	"sort"
	"sync"
)

var _ Balancer = &ConsistentHash{}

type ConsistentHash struct {
	replicate int
	hashFunc  func(v []byte) uint32
	ring      []*ringNode
	ringBuf   []*ringNode
	valMap    map[Node]struct{}
	lock      *sync.RWMutex
}

type ringNode struct {
	hash uint32
	val  Node
}

func NewConsistentHash(rep int, hashFunc func([]byte) uint32) *ConsistentHash {
	ch := &ConsistentHash{
		replicate: rep,
		hashFunc:  hashFunc,
		valMap:    make(map[Node]struct{}),
		ringBuf:   make([]*ringNode, rep),
		lock:      &sync.RWMutex{},
	}
	return ch
}

func (ch *ConsistentHash) Add(nodes ...Node) (err error) {
	ch.lock.Lock()
	length := len(ch.ring)
	defer func() {
		if len(ch.ring) != length {
			ch.sortRing()
		}
		ch.lock.Unlock()
	}()
	for _, n := range nodes {
		if _, ok := ch.valMap[n]; ok {
			continue
		}
		for i := 0; i < ch.replicate; i++ {
			str := fmt.Sprintf("%s#%d", n.String(), i)
			ch.ringBuf[i] = &ringNode{
				hash: ch.hashFunc([]byte(str)),
				val:  n,
			}
		}
		ch.ring = append(ch.ring, ch.ringBuf...)
		ch.resetRingBuf()
		ch.valMap[n] = struct{}{}
	}
	return
}

func (ch *ConsistentHash) Delete(nodes ...Node) (err error) {
	ch.lock.Lock()
	length := len(ch.ring)
	defer func() {
		if length != len(ch.ring) {
			ch.sortRing()
		}
		ch.lock.Unlock()
	}()
	for _, n := range nodes {
		if _, ok := ch.valMap[n]; !ok {
			continue
		}
		for i := 0; i < ch.replicate; i++ {
			length = len(ch.ring)
			objHash := ch.hashFunc([]byte(fmt.Sprintf("%s#%d", n.String(), i)))
			target := sort.Search(length, func(k int) bool {
				return ch.ring[k].hash >= objHash
			})
			ch.ring = append(ch.ring[:target], ch.ring[target+1:]...)
		}
		delete(ch.valMap, n)
	}
	return
}

func (ch *ConsistentHash) GetNode(pkg []byte) (res Node, ok bool) {
	ch.lock.RLock()
	defer ch.lock.RUnlock()
	length := len(ch.ring)
	objHash := ch.hashFunc(pkg)
	target := sort.Search(length, func(i int) bool {
		return ch.ring[i].hash >= objHash
	})
	res = ch.ring[target%length].val
	ok = true
	return
}

func (ch *ConsistentHash) sortRing() {
	sort.Slice(ch.ring, func(i, j int) bool {
		return ch.ring[i].hash < ch.ring[j].hash
	})
}

func (ch *ConsistentHash) resetRingBuf() {
	for i := 0; i < ch.replicate; i++ {
		ch.ringBuf[i] = nil
	}
}

func (ch *ConsistentHash) Show() {
	for _, v := range ch.ring {
		fmt.Println(v.hash, "\t", v.val.String())
	}
}
