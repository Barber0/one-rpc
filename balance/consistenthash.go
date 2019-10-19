package balance

var _ Balancer = &ConsistentHash{}

type ConsistentHash struct {
	replicate	int32
	hashFunc	func([]byte) uint32
}

func NewConsistentHash(rep int32, hashFunc func([]byte) uint32) *ConsistentHash {
	ch := &ConsistentHash{
		replicate:		rep,
		hashFunc:		hashFunc,
	}
	return ch
}

func (ch *ConsistentHash) Add(nodes ...Node) (err error) {
	return
}

func (ch *ConsistentHash) Delete(nodes ...Node) (err error) {
	return
}

func (ch *ConsistentHash) GetNode(v []byte) (res Node, ok bool) {
	return
}