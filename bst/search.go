package bst

func (bst *BST) Search(value int) *Node {
	return search(bst.Root, value)
}

func search(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value == node.Value {
		return node
	} else if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

func (bst *BST) SearchMin() *Node {
	current := bst.Root
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BST) SearchMax() *Node {
	current := bst.Root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
