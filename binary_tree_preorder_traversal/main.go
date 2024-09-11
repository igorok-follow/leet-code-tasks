package main

import "log"

func main() {
	log.Println(preorderTraversal(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  4,
					Left: nil,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  6,
						Left: nil,
					},
					Right: &TreeNode{
						Val:  7,
						Left: nil,
					},
				},
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:  9,
						Left: nil,
					},
				},
			},
		},
	))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var (
		stack   = make([]*TreeNode, 0)
		res     = make([]int, 0)
		current *TreeNode
	)

	if root == nil {
		return res
	}

	stack = append(stack, root)

	for len(stack) != 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, current.Val)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return res
}
