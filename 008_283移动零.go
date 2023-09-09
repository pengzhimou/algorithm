package main

import "fmt"

// func main() {
// 	// moveZero([]int{1, 0, 0, 2, 3, 4})
// 	// moveZeroes([]int{1, 0, 0, 2, 3, 4})
// 	moveZeroes([]int{1, 2, 0, 2, 3, 4})
// }

func moveZero(array []int) { //移动的方式
	j := 0
	temp := 0
	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			array[j] = array[i]
			j++
		} else {
			temp++ // 计算后面需要覆盖0的个数
		}
	}
	fmt.Println(array)

	for idx := len(array) - temp; idx < len(array); idx++ {
		array[idx] = 0
	}
	fmt.Println(array)
}

func moveZeroes(nums []int) { //交换的方式
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	fmt.Println(nums)
}
