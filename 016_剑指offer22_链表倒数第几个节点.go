package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthNode(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	for k > 1 && tail != nil {
		tail = tail.Next
		k--
	}
	if tail == nil {
		return nil
	}
	for tail.Next != nil {
		tail = tail.Next
		head = head.Next
	}
	return head
}

func ssss() {
	//test
	l3 := &ListNode{3, nil}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}
	fmt.Println("1 -> 2 -> 3")

	fmt.Println("reverse 1th")
	fmt.Println(kthNode(l1, 1))
	fmt.Println("reverse 2th")
	fmt.Println(kthNode(l1, 2))
	fmt.Println("reverse 3th")
	fmt.Println(kthNode(l1, 3))
	fmt.Println("reverse 4th")
	fmt.Println(kthNode(l1, 4))

}

// https://github.com/DinghaoLI/Coding-Interviews-Golang

// 1，反转链表
// 2，正数个数-目标个数+1

// 1，遍历链表，到结尾得到总长度
// 2.继续遍历 正数个数-目标个数+1
