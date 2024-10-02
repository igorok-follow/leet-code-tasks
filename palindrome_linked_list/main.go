package main

import "log"

func main() {
	log.Println(isPalindrome(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	if head.Next == nil {
		return true
	}

	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverseList(slow)

	for head != nil && slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head, slow = head.Next, slow.Next
	}

	return true
}

func reverseList(node *ListNode) *ListNode {
	var previous, current *ListNode = nil, node
	for current != nil {
		previous, current, current.Next = current, current.Next, previous
	}
	return previous
}
