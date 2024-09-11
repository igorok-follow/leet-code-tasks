package main

import "log"

var intersect = &ListNode{
	Val: 8,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	},
}

func main() {
	a := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: intersect,
		},
	}
	b := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: intersect,
			},
		},
	}

	log.Println(getIntersectionNode(
		a,
		b,
	))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)

	n := headA
	for n != nil {
		seen[n] = true
		n = n.Next
	}

	n = headB
	for n != nil {
		if seen[n] {
			return n
		}
		n = n.Next
	}

	return nil
}
