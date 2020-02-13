package common

import (
	"fmt"
	"testing"
)

func Test_BinarySearchInt(t *testing.T) {
	fmt.Println(BinarySearchFirstGT([]int{1, 2, 3}, 2))
	fmt.Println(BinarySearchFirstGT([]int{1, 2, 4}, 3))
	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4}, 2))

	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4}, 1))
	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4}, 4))

	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4}, 100))
	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4}, 0))

	fmt.Println(BinarySearchFirstGT([]int{1, 2, 3, 5}, 2))
	fmt.Println(BinarySearchFirstGT([]int{1, 2, 4, 6}, 3))
	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4, 6}, 2))
	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4, 6}, 1))
	fmt.Println(BinarySearchFirstGT([]int{1, 3, 4, 6}, 4))
}