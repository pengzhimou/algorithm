// https: //blog.csdn.net/weixin_42324313/article/details/121955235
// 快速排序，堆排序
package main

import "fmt"

// 选择排序 每一轮把本轮最小的放到本轮最前面互换位置
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minNumIndex := i // 无序区第一个
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minNumIndex] {
				minNumIndex = j
			}
		}
		arr[i], arr[minNumIndex] = arr[minNumIndex], arr[i]
	}
	fmt.Println(arr)
}

func main123dg() {
	aa := []int{5, 4, 3, 2, 1}
	// BubbleSort2(aa)
	fmt.Println(aa)
}
