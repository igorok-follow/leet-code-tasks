package main

import "log"

func main() {
	tree := invertTree(
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
	)

	log.Println(tree.Left.Left, tree.Left.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	if root.Right != nil {
		root.Right = invertTree(root.Right)
	}
	if root.Left != nil {
		root.Left = invertTree(root.Left)
	}

	return root
}

func iterative(root *TreeNode) []int {
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
