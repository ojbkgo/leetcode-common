package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array2List(data []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	if len(data) == 0 {
		return nil
	}

	head = &ListNode{
		Val: data[0],
	}

	if len(data) == 1 {
		return head
	}

	tail = head

	for _, v := range data[1:] {
		node := &ListNode{
			Val:v,
		}

		tail.Next = node
		tail = node
	}

	return head
}

func TravelList(list *ListNode, action func(v int)) {
	for list != nil {
		action(list.Val)
		list = list.Next
	}
}

func DumpList(list *ListNode) {
	data := make([]int, 0)

	TravelList(list, func(v int) {
		data = append(data, v)
	})

	Dump(data)
}