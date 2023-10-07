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
		// 削除対象のノードが見つかったとき

		if node.Left == nil && node.Right == nil {
			return nil
		}
		// 削除対象のノードが子ノードを1つだけ持つ場合、その子ノードを削除対象のノードにすげ替える
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// 削除対象のノードが子ノードを2つ持つ場合、右側サブツリーの最小値を削除対象のノードにすげ替える
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
