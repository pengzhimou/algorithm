package main

import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func main() {

// 	l := &ListNode{3, &ListNode{4, &ListNode{5, nil}}}
// 	l.Next.Next.Next = l

// 	a := &ListNode{1, &ListNode{2, l}}

// 	fmt.Println(
// 		a.Val,
// 		a.Next.Val,
// 		a.Next.Next.Val,
// 		a.Next.Next.Next.Val,
// 		a.Next.Next.Next.Next.Val,
// 		a.Next.Next.Next.Next.Next.Val,
// 		a.Next.Next.Next.Next.Next.Next.Val,
// 		a.Next.Next.Next.Next.Next.Next.Next.Val,
// 		a.Next.Next.Next.Next.Next.Next.Next.Next.Val,
// 	)

// 	detectCycle2(
// 		a,
// 	)

// }

//链表里是否有环，并给出有环的第一项

func detectCycle1(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			fmt.Println(1111, fast, slow, fast.Val, slow.Val)
			p := head
			for p != slow { // 追击由于会追上，会夺走一个slow，所以要回退，从头遍历一次到slow直接return
				p = p.Next
				slow = slow.Next
			}

			fmt.Println(2222, p, p.Val)
			return p
		}
	}
	return nil
}

// 1 2 3 4 5 3 4 5 3
// 1111 &{4 0xc000108070} &{4 0xc000108070} 4 4
// 2222 &{3 0xc000108060} 3
