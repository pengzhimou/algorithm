package main

import "fmt"

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

func main() {
	x := TreeNode3{
		Val: 1,
		Left: &TreeNode3{
			Val: 2,
			Left: &TreeNode3{
				Val: 4,
				Left: &TreeNode3{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode3{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode3{
			Val: 3,
			Left: &TreeNode3{
				Val:  7,
				Left: nil,
				Right: &TreeNode3{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode3{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	_ = x

	aa := postorderTraversal1(&x)
	fmt.Println(aa)
	// [5 6 4 2 9 7 8 3 1]
}
func postorderTraversal1(root *TreeNode3) (res []int) {
	var postorder func(*TreeNode3)
	postorder = func(node *TreeNode3) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

func postorderTraversal2(root *TreeNode3) (res []int) {
	stk := []*TreeNode3{}
	var prev *TreeNode3
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func postorderTraversal3(root *TreeNode3) (res []int) {
	addPath := func(node *TreeNode3) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}
