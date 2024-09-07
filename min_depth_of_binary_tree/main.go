package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	x := minDepth(root.Left)
	y := minDepth(root.Right)

	if x > 0 || y > 0 {
		if min(x, y) == 0 {
			return max(x, y) + 1
		}

		return min(x, y) + 1
	}

	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
