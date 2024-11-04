package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	log.Println(diameterOfBinaryTree(&TreeNode{
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
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}))
}

var m int

func diameterOfBinaryTree(root *TreeNode) int {
	m = 0
	dfs(root)

	return m
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthL := dfs(root.Left)
	depthR := dfs(root.Right)

	if depthL+depthR > m {
		m = depthL + depthR
	}
	log.Println(root.Val, m)

	return max(depthL, depthR) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
