package bst

type BinarySearchTree struct {
	root *Node
}

func (binaryTree *BinarySearchTree) Init(root *Node) {
	binaryTree.root = root
}

func (binaryTree *BinarySearchTree) Add(value int) *Node {
	return addHelper(binaryTree.root, value)
}

func addHelper(current *Node, value int) *Node {
	if current == nil {
		n := new(Node)
		n.Init(value)
		return n
	}

	if value < current.value {
		current.left = addHelper(current.left, value)
	} else if value > current.value {
		current.right = addHelper(current.right, value)
	} else {
		return current
	}
	
	return current
}

func (binaryTree *BinarySearchTree) Contains(value int) bool {
	return containsHelper(binaryTree.root, value)
}

func containsHelper(current *Node, value int) bool {
	if current == nil {
		return false
	}
	if value == current.value {
		return true
	}
	if value < current.value {
		return containsHelper(current.left, value)
	} else {
		return containsHelper(current.right, value)
	}
}

func (binaryTree *BinarySearchTree) Delete(value int) {
	deleteHelper(binaryTree.root, value)
}

func deleteHelper(current *Node, value int) *Node {
	if current == nil {
		return nil
	}
	if value == current.value {
		if current.right == nil && current.left == nil {
			return nil
		}
		if current.right == nil {
			return nil
		}
		if current.left == nil {
			return nil
		}
		smallestValue := findSmallestValue(current.right)
		current.value = smallestValue
		current.right = deleteHelper(current.right, smallestValue)
		return current
	}
	if value < current.value {
		current.left = deleteHelper(current.left, value)
		return current
	}
	current.right = deleteHelper(current.right, value)
	return current
}

func findSmallestValue(root *Node) int {
	if root.left == nil {
		return root.value
	} else {
		return findSmallestValue(root.left)
	}
}

func TraverseInOrder(root *Node) {
	if root != nil {
		TraverseInOrder(root.left)
		print(" " + string(root.value))
		TraverseInOrder(root.right)
	}
}

func createTestBinaryTree() *BinarySearchTree {
	bt := new(BinarySearchTree)
	n := new(Node)
	n.Init(6)
	bt.Init(n)
	bt.Add(6)
	bt.Add(4)
	bt.Add(8)
	bt.Add(3)
	bt.Add(5)
	bt.Add(7)
	bt.Add(9)
	return bt
}