package main

import (
	"log"
	"time"
)

var h = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	},
}

func main() {
	t := time.Now()
	v := deleteDuplicates(h)
	log.Println(time.Since(t), "\n")

	c := v
	for c != nil {
		log.Println(c.Val)
		c = c.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Val != current.Next.Val {
			current = current.Next
		} else {
			current.Next = current.Next.Next
		}
	}

	return head
}
