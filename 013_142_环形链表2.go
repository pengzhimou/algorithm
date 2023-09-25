package main

// func main() {
// 	aa := mergeTwoLists2(
// 		&ListNode{2, &ListNode{4, &ListNode{5, &ListNode{5, nil}}}},
// 		&ListNode{2, &ListNode{3, &ListNode{4, nil}}},
// 	)

//链表里是否有环，并给出有环的第一项

func detectCycle(head *ListNode) *ListNode {
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
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
