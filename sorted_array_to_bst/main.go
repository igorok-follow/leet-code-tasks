package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})

	log.Println(root.Left, root.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	m := len(nums) / 2

	if len(nums) == 0 {
		return nil
	}

	node := &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}

	return node
}
