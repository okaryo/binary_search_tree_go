package bst

import "testing"

func TestSearch(t *testing.T) {
	bst := New()
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	// 存在する値に対するテスト
	for _, v := range values {
		if node := bst.Search(v); node == nil || node.Value != v {
			t.Errorf("Expected to find node with value %d in the BST, but did not", v)
		}
	}

	// 存在しない値に対するテスト
	nonExistentValues := []int{1, 9, 10, 0, -5}
	for _, v := range nonExistentValues {
		if node := bst.Search(v); node != nil {
			t.Errorf("Expected value %d to not exist in the BST, but found node with value %d", v, node.Value)
		}
	}
}

func TestMinAndMax(t *testing.T) {
	bst := New()
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	if minNode := bst.SearchMin(); minNode == nil || minNode.Value != 2 {
		t.Errorf("Expected Min node to have value 2, but got %v", minNode)
	}

	if maxNode := bst.SearchMax(); maxNode == nil || maxNode.Value != 8 {
		t.Errorf("Expected Max node to have value 8, but got %v", maxNode)
	}
}
