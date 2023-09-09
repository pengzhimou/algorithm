package main

import "fmt"

// func main() {

// 	// fmt.Println(merge1(
// 	// 	[]int{1, 2, 4}, 3,
// 	// 	[]int{2, 3, 5, 6}, 4,
// 	// ))

// 	a := make([]int, 3, 7)
// 	a[0] = 1
// 	a[1] = 2
// 	a[2] = 4

// 	merge1(
// 		a, 3,
// 		[]int{2, 3, 5, 6, 6}, 5,
// 	)
// }

func merge1(nums1 []int, m int, nums2 []int, n int) []int {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	// copy(nums1, sorted) // 原slice只有长度5，所以会丢弃剩余的值
	// return nums1

	return sorted
}

func merge3(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
	fmt.Println(nums1)
}
