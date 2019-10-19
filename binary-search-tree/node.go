package bst

type Node struct {
	value int
	left *Node
	right *Node
}

func (n *Node) Init(value int) {
	n.value = value
}