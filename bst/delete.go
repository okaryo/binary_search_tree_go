package bst

func (bst *BST) Delete(value int) {
	bst.Root = deleteNode(bst.Root, value)
}

func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		// When the target node for deletion is found

		// If the target node has no children
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// If the target node has only one child, replace the target node with its child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// If the target node has two children, replace the target node with the smallest value in the right subtree
		minLargest := findMinValue(node.Right)
		node.Value = minLargest
		node.Right = deleteNode(node.Right, minLargest)
	}

	return node
}

func findMinValue(node *Node) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}
