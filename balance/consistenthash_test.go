package balance_test

import (
	"fmt"
	"github.com/Barber0/one-rpc/balance"
	"hash/crc32"
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

	ch := balance.NewConsistentHash(2,crc32.ChecksumIEEE)
	ch.Add(a,b,c)
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
