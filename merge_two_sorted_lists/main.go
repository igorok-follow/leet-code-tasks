package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	sorted := mergeTwoLists(list1, list2)
	for sorted != nil {
		log.Println(sorted.Val)
		sorted = sorted.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := make([]int, 0, 0)
	for list1 != nil {
		list = append(list, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		list = append(list, list2.Val)
		list2 = list2.Next
	}

	for {
		ready := true
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				ready = false
			}
		}
		if ready {
			break
		}
	}

	var (
		parentNode *ListNode
		cursor     *ListNode
	)

	for i := 0; i < len(list); i++ {
		current := &ListNode{
			Val:  list[i],
			Next: nil,
		}

		if i == 0 {
			parentNode = current
			cursor = parentNode
		} else {
			cursor.Next = current
			cursor = cursor.Next
		}
	}

	return parentNode
}
