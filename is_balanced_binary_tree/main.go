package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	isB, _ := dfs(root)
	return isB
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isL, depthL := dfs(root.Left)
	isR, depthR := dfs(root.Right)

	if isL && isR && abs(depthL-depthR) <= 1 {
		return true, max(depthL, depthR) + 1
	}

	return false, 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
