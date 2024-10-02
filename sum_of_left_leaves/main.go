package main

import (
	"log"
)

func main() {
	log.Println(sumOfLeftLeaves(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	dfs(root, &sum)
	return sum
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			*sum += root.Left.Val
		}
		dfs(root.Left, sum)
	}

	if root.Right != nil {
		dfs(root.Right, sum)
	}
}
