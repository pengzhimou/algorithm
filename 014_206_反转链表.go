package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	a := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	b := reverseList(a)

	fmt.Println(
		b.Val,
		b.Next.Val,
		b.Next.Next.Val,
		b.Next.Next.Next.Val,
		b.Next.Next.Next.Next.Val,
	)
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}
