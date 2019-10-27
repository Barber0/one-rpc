package balance_test

import (
	"fmt"
	"github.com/Barber0/one-rpc/balance"
	"hash/crc32"
	"math/rand"
	"sync"
	"testing"
)

type N struct {
	v string
}

func (n *N) String() string {
	return n.v
}

func TestNewConsistentHash(t *testing.T) {
	a := &N{"alpha"}
	b := &N{"beta"}
	c := &N{"delta"}

	ch := balance.NewConsistentHash(2, crc32.ChecksumIEEE)
	ch.Add(a, b, c)
	ch.Show()

	fmt.Println("-------------")

	ch.Delete(b)
	ch.Show()

	fmt.Println("+++++++++++")

	fmt.Println(ch.GetNode([]byte("dddd")))

	//a := []int{1,2,3,4,5}
	//fmt.Println(sort.Search(len(a), func(i int) bool {
	//	return a[i] >= 4
	//}))
}

var (
	rep     = 32
	hash    = crc32.ChecksumIEEE
	testLen = 10
	nodes   = 10
	sams    = 20
	tpl     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func BenchmarkNewConsistentHash1(b *testing.B) {
	ch := balance.NewConsistentHash(rep, hash)
	nds := getTestNodes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Add(nds...)
		ch.Delete(nds...)
	}
}

func getTestNodes() []balance.Node {
	tns := make([]balance.Node, nodes)
	idx := 0
	defer func() {
		if pa := recover(); pa != nil {
			fmt.Println("ddddddddd\t", idx)
		}
	}()
	for i := 0; i < nodes; i++ {
		tn := &testNode{}
		for j := 0; j < testLen; j++ {
			idx = int(rand.Uint32()) % len(tpl)
			tn.v += tpl[idx : idx+1]
		}
		tns[i] = tn
	}
	return tns
}

type testNode struct {
	v string
}

func (t *testNode) String() string {
	return t.v
}

func TestLock(t *testing.T) {
	a := &sync.RWMutex{}
	a.Lock()
	defer a.Unlock()
	fmt.Println("fffff")
}
