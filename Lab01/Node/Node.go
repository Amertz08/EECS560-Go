package Node

type Node struct {
	Value int
	Next *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}
