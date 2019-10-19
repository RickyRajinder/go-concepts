package bst

type BinarySearchTree struct {
	root *Node
}

func addHelper(current *Node, value int) *Node {
	if current == nil {
		n := new(Node)
		n.Init(value)
		return n
	}

}