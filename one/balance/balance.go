package balance

type Balancer interface {
	Add(nodes ...Node) error
	Delete(nodes ...Node) error
	GetNode(v []byte) (Node, bool)
}

type Node interface {
	String() string
}