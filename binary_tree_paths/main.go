package main

import (
	"fmt"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	history := make([]string, 0)
	dfs(root, history, &res)

	return res
}

func dfs(root *TreeNode, history []string, res *[]string) {
	if root == nil {
		return
	}

	history = append(history, fmt.Sprintf("%d", root.Val))

	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(history, "->"))
		return
	}

	dfs(root.Left, history, res)
	dfs(root.Right, history, res)
}
