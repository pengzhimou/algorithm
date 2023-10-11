package main

//所有子树 高度差在1

type TreeNode6 struct {
	Val   int
	Left  *TreeNode6
	Right *TreeNode6
}

func isBalanced(root *TreeNode6) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode6) int {
	if root == nil {
		return 0
	}
	return max1(height(root.Left), height(root.Right)) + 1
}

func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

/////////////////////////

func isBalanced2(root *TreeNode6) bool {
	return height2(root) >= 0
}

func height2(root *TreeNode6) int {
	if root == nil {
		return 0
	}
	leftHeight := height2(root.Left)
	rightHeight := height2(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs2(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max2(leftHeight, rightHeight) + 1
}

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs2(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
