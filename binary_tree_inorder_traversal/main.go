package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	log.Println(inorderTraversal(
		&TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
	))
	log.Println(time.Since(t), "\n")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		stack   = make([]*TreeNode, 0)
		current *TreeNode
		res     = make([]int, 0)
	)

	current = root
	for len(stack) != 0 || current != nil {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, current.Val)
		current = current.Right
	}

	return res
}
