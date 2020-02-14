package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DumpTree(tree *TreeNode) {
	fmt.Println("pre:", TravelTreePre(tree))
	fmt.Println("mid:", TravelTreeMid(tree))
}

func Array2Tree(data []int) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	plain := make([]*TreeNode, 0)

	for _, item := range data {
		if item > 0 {
			plain = append(plain, &TreeNode{Val:item})
		} else {
			plain = append(plain, nil)
		}
	}

	for i := 1; i < len(plain); i++ {
		if plain[i] != nil {
			var pidx int
			if i % 2 == 1 {
				pidx = i / 2
				plain[pidx].Left = plain[i]
			} else {
				pidx = i / 2 - 1
				plain[pidx].Right = plain[i]
			}
		}
	}

	return plain[0]
}


func TravelTreePre(root *TreeNode) []int {
	seq := make([]int, 0)

	_travelPre(root, func(v int) {
		seq = append(seq, v)
	})

	return seq
}

func _travelPre(root *TreeNode, fn func(v int)) {
	if root == nil {
		return
	}

	fn(root.Val)

	_travelPre(root.Left, fn)
	_travelPre(root.Right, fn)
}


func TravelTreeMid(root *TreeNode) []int {
	seq := make([]int, 0)

	_travelMid(root, func(v int) {
		seq = append(seq, v)
	})

	return seq
}

func _travelMid(root *TreeNode, fn func(v int)) {
	if root == nil {
		return
	}

	_travelMid(root.Left, fn)
	fn(root.Val)
	_travelMid(root.Right, fn)
}
