package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	if isL := hasPathSum(root.Left, targetSum); isL {
		return true
	}
	if isR := hasPathSum(root.Right, targetSum); isR {
		return true
	}

	return false
}
