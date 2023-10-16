// https: //blog.csdn.net/weixin_42324313/article/details/121955235
// 快速排序，堆排序
package main

import "fmt"

// 插入排序 易于理解一些
func InsertionSort1(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		tmp := arr[i]
		j := i - 1
		for ; j >= 0 && tmp < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = tmp
	}
	fmt.Println(arr)
}

// 改进版插入排序
func InsertionSort2(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ { // 无序区
		tmp := arr[i]
		left, right := 0, i-1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] > tmp {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		j := i - 1
		for ; j >= left; j-- { // 有序区
			arr[j+1] = arr[j]
		}
		arr[left] = tmp
	}
	fmt.Println(arr)
}

func main123123() {
	// aa := []int{5, 4, 3, 2, 1}
	aa := []int{5, 4, 1, 2, 3}
	InsertionSort1(aa)
	fmt.Println(aa)

	for i := range aa {
		fmt.Println(i)
	}
}
