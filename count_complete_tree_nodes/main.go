package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nleft := 1
	nright := 1

	l := root.Left
	r := root.Right
	for l != nil {
		nleft++
		l = l.Left
	}

	for r != nil {
		nright++
		r = r.Left
	}

	if nleft == nright {
		return 1<<nleft - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
