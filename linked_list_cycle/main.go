package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next

		if head == fast {
			return true
		}
	}

	return false
}
