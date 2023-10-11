package main

import "fmt"

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func mainfff() {

	x := TreeNode2{
		Val: 1,
		Left: &TreeNode2{
			Val: 2,
			Left: &TreeNode2{
				Val: 4,
				Left: &TreeNode2{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode2{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode2{
			Val: 3,
			Left: &TreeNode2{
				Val:  7,
				Left: nil,
				Right: &TreeNode2{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode2{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	_ = x

	aaf := preorderTraversal1(&x)

	fmt.Println(aaf)
	// [1 2 4 5 6 3 7 9 8]
}

func preorderTraversal1(root *TreeNode2) (vals []int) {
	var preorder func(*TreeNode2)
	preorder = func(node *TreeNode2) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

func preorderTraversal2(root *TreeNode2) (vals []int) {
	stack := []*TreeNode2{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func preorderTraversal3(root *TreeNode2) (vals []int) {
	var p1, p2 *TreeNode2 = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
	return
}
