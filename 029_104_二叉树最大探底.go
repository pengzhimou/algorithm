package main

type TreeNode5 struct {
	Val   int
	Left  *TreeNode5
	Right *TreeNode5
}

func maxDepth(root *TreeNode5) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
