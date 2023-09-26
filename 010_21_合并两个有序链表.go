package main

// func main() {
// 	aa := mergeTwoLists2(
// 		&ListNode{2, &ListNode{4, &ListNode{5, &ListNode{5, nil}}}},
// 		&ListNode{2, &ListNode{3, &ListNode{4, nil}}},
// 	)

// 	fmt.Println(aa)
// 	fmt.Println(aa.Next)
// 	fmt.Println(aa.Next.Next)
// 	fmt.Println(aa.Next.Next.Next)
// 	fmt.Println(aa.Next.Next.Next.Next)
// 	fmt.Println(aa.Next.Next.Next.Next.Next)
// 	fmt.Println(aa.Next.Next.Next.Next.Next.Next)
// 	fmt.Println(aa.Next.Next.Next.Next.Next.Next.Next)

// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists1(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 用哨兵节点简化代码逻辑
	cur := dummy         // cur 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2 // 注：如果都为空则返回空，递归时，替代拼接剩余链表
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}
