package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	bst := New()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)

	if bst.Root.Value != 5 {
		t.Errorf("Expected root value to be 5, but got %d", bst.Root.Value)
	}
	if bst.Root.Left.Value != 3 {
		t.Errorf("Expected left child value to be 3, but got %d", bst.Root.Left.Value)
	}
	if bst.Root.Right.Value != 7 {
		t.Errorf("Expected right child value to be 7, but got %d", bst.Root.Right.Value)
	}
}
