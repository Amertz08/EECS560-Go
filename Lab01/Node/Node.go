package Node

type Node struct {
	Value int
	Next *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

func (n *Node) GetValue() int {
	return n.Value
}

func (n *Node) setValue(value int) {
	n.Value = value
}