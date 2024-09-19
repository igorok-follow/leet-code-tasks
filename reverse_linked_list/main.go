package main

import "log"

func main() {
	head := reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	})

	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	head, stack = stack[len(stack)-1], stack[:len(stack)-1]
	var (
		next    *ListNode
		current = head
	)
	for {
		next, stack = stack[len(stack)-1], stack[:len(stack)-1]

		current.Next = next
		current = next

		if len(stack) == 0 {
			current.Next = nil
			break
		}
	}

	return head
}
