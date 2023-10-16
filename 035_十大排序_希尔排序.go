// https: //blog.csdn.net/weixin_42324313/article/details/121955235
// 快速排序，堆排序
package main

import "fmt"

// 希尔排序
func ShellSort(arr []int) {
	n := len(arr)
	for gap := n / 2; gap >= 1; gap = gap / 2 { // 缩小增量序列，希尔建议每次缩小一半
		for i := gap; i < n; i++ { // 子序列
			tmp := arr[i]
			j := i - gap
			for ; j >= 0 && tmp < arr[j]; j = j - gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = tmp
		}
	}
	fmt.Println(arr)
}

func main() {
	// aa := []int{5, 4, 3, 2, 1}
	aa := []int{5, 4, 1, 2, 3}
	InsertionSort1(aa)
	fmt.Println(aa)

	for i := range aa {
		fmt.Println(i)
	}
}
