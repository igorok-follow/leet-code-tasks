package main

import "log"

func main() {
	log.Println(isSameTree(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil || p == nil {
		return q == p
	}

	return q.Val == p.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
