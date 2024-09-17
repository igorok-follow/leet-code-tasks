package main

import "log"

func main() {
	head := removeElements(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}, 2)

	current := head
	for current != nil {
		log.Println(current.Val)
		current = current.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	root := head
	current := root
	var prev *ListNode

	for current != nil {
		if current.Val == val && current == root {
			root = current.Next
		} else if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}

		current = current.Next
	}

	return root
}
