package main

import "fmt"

func main() {
	x := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}

	xxx := inorderTraversal1(&x)
	fmt.Println(xxx)

	yyy := inorderTraversal2(&x)
	fmt.Println(yyy)

	// [5 4 6 2 1 7 9 3 8]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(tn *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder =
		func(node *TreeNode) {
			if node == nil {
				return
			}
			inorder(node.Left)
			res = append(res, node.Val)
			inorder(node.Right)
		}
	inorder(tn)
	return
}

func inorderTraversal2(tn *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for tn != nil || len(stack) > 0 {
		for tn != nil {
			stack = append(stack, tn)
			tn = tn.Left
		}
		tn = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, tn.Val)
		tn = tn.Right
	}
	return
}
