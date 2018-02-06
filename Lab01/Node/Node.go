package Node

type Node struct {
	value int
	next *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

func (n *Node) getValue() int {
	return n.value
}

func (n *Node) setValue(value int) {
	n.value = value
}