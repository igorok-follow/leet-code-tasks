package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseStr(s string, k int) string {
	l, r := 0, k-1

	b := []byte(s)
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	return string(b)
}

func main() {
	log.Println('Z')

	//nums := addTwoNumbers(
	//	&ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 4,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: nil,
	//			},
	//		},
	//	},
	//	&ListNode{
	//		Val: 5,
	//		Next: &ListNode{
	//			Val: 6,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: nil,
	//			},
	//		},
	//	},
	//)
	//
	//for nums != nil {
	//	log.Println(nums.Val)
	//	nums = nums.Next
	//}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := rec(l1, 1) + rec(l2, 1)
	head := &ListNode{
		Val:  res % 10,
		Next: nil,
	}
	res /= 10

	current := head
	for res != 0 {
		current.Next = &ListNode{
			Val:  res % 10,
			Next: nil,
		}
		res /= 10
		current = current.Next
	}

	return head
}

func rec(node *ListNode, mult int) int {
	if node == nil {
		return 0
	}

	num := rec(node.Next, mult*10)

	return num + node.Val*mult
}
