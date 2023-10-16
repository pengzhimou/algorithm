// https: //blog.csdn.net/weixin_42324313/article/details/121955235
// 快速排序，堆排序
package main

import "fmt"

// 冒泡排序 每一轮把当前找到最小的放到本轮的第一位(i)，从而最终变成最小到大排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

// 改进的冒泡排序
func BubbleSort2(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ { // j<i，一个增加一个减小，可以减少迭代次数，因为j轮每一次已经把本轮最大的放到队尾了，每轮可以少迭代一些
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// for i := 0; i < n-1; i++ {
	// 	for j := 0; j < n-1; j++ {
	// 		if arr[j] > arr[j+1] {
	// 			arr[j], arr[j+1] = arr[j+1], arr[j]
	// 		}
	// 	}
	// }
	fmt.Println(arr)
}

func main12312d() {
	aa := []int{5, 4, 3, 2, 1}
	// BubbleSort2(aa)
	fmt.Println(aa)
}
