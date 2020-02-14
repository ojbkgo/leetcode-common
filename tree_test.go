package common

import "testing"

func Test_Tree(t *testing.T) {
	tree := []int{1, 2, 3, 0, 4, 5, 0, 0, 0, 0, 0, 6, 7}

	DumpTree(Array2Tree(tree))
}
