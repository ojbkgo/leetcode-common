package common

import "testing"

func Test_Array2List(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	DumpList(Array2List(data))
	data = []int{1, 2}
	DumpList(Array2List(data))
	data = []int{1}
	DumpList(Array2List(data))
	data = []int{}
	DumpList(Array2List(data))
}