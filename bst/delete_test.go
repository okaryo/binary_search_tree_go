package bst

import "testing"

func TestDelete(t *testing.T) {
	tests := []struct {
		initValues  []int
		deleteValue int
		expectRoot  *Node
	}{
		{
			initValues:  []int{5},
			deleteValue: 5,
			expectRoot:  nil,
		},
		{
			initValues:  []int{5, 3},
			deleteValue: 5,
			expectRoot:  &Node{Value: 3},
		},
		{
			initValues:  []int{5, 3, 7},
			deleteValue: 5,
			expectRoot:  &Node{Value: 7, Left: &Node{Value: 3}},
		},
		{
			initValues:  []int{5, 3, 7, 2, 4, 6, 8},
			deleteValue: 3,
			expectRoot:  &Node{Value: 5, Left: &Node{Value: 4, Left: &Node{Value: 2}}, Right: &Node{Value: 7, Left: &Node{Value: 6}, Right: &Node{Value: 8}}},
		},
	}

	for _, test := range tests {
		bst := New()
		for _, v := range test.initValues {
			bst.Insert(v)
		}
		bst.Delete(test.deleteValue)
		if !compareTrees(bst.Root, test.expectRoot) {
			t.Errorf("After deleting %d, expected BST root to be %v, but got %v", test.deleteValue, test.expectRoot, bst.Root)
		}
	}
}

func compareTrees(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Value == b.Value && compareTrees(a.Left, b.Left) && compareTrees(a.Right, b.Right)
}
