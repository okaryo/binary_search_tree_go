package bst

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func New() *BST {
	return &BST{}
}
