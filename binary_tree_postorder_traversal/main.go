package main

import (
	"log"
)

func main() {
	log.Println(postorderTraversal(&TreeNode{
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

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	dfs(root.Right, res)

	*res = append(*res, root.Val)
}

func iterative(root *TreeNode) []int {
	var res = make([]int, 0)
	if root == nil {
		return res
	}

	var (
		current *TreeNode
		stack   = make([]*TreeNode, 0)
	)
	stack = append(stack, root)

	for len(stack) != 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append([]int{current.Val}, res...)

		if current.Left != nil {
			stack = append(stack, current.Left)
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return res
}
