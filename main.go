package main

import (
	"bst/bst"
	"fmt"
)

func main() {
	bst := bst.New()

	// Insert
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	// Search
	searchValues := []int{5, 3, 9, 1}
	for _, v := range searchValues {
		if node := bst.Search(v); node != nil {
			fmt.Printf("%d exists in the BST.\n", v)
		} else {
			fmt.Printf("%d does not exist in the BST.\n", v)
		}
	}

	// SearchMin, SearchMax
	fmt.Printf("Minimum value in the BST is: %d\n", bst.SearchMin().Value)
	fmt.Printf("Maximum value in the BST is: %d\n", bst.SearchMax().Value)

	// Delete
	deleteValues := []int{3, 5}
	for _, v := range deleteValues {
		bst.Delete(v)
		if bst.Search(v) != nil {
			fmt.Printf("Failed to delete %d from the BST.\n", v)
		} else {
			fmt.Printf("Successfully deleted %d from the BST.\n", v)
		}
	}
}
