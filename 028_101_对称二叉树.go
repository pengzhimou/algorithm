package main

type TreeNode4 struct {
	Val   int
	Left  *TreeNode4
	Right *TreeNode4
}

func isSymmetric(root *TreeNode4) bool {
	return check(root, root)
}

func check(p, q *TreeNode4) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
