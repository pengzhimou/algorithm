package main

import "fmt"

func mainc() {
	aa := MergeSort1([]int{7, 6, 5, 4, 3, 2, 1})
	fmt.Println(aa)
}

// 归并排序--递归版
func MergeSort1(arr []int) []int {
	n := len(arr)
	if n < 2 { // 健壮性
		return arr
	}
	mid := n / 2
	left := arr[:mid]
	right := arr[mid:]
	return mergeX1(MergeSort1(left), MergeSort1(right))
}

func mergeX1(left, right []int) []int {
	res := []int{}
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:] // 将头一个直接切出去
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) == 0 { // left结束，right剩下的直接拖下来
		res = append(res, right...)
	}
	if len(right) == 0 { // right结束，left剩下的直接拖下来
		res = append(res, left...)
	}
	return res
}

/////////////////////////

// 归并排序--迭代版
func MergeSort2(arr []int) []int {
	n := len(arr)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for step := 1; step <= n; step <<= 1 { // 外层控制步长
		offset := 2 * step
		for i := 0; i < n; i += offset { // 内层控制分组
			h2 := min(i+step, n-1)        // 第二段头部，防止超过数组长度
			tail2 := min(i+offset-1, n-1) // 第二段尾部
			mergeX2(arr, i, h2, tail2)
		}
	}
	return arr
}

func mergeX2(arr []int, h1 int, h2 int, tail2 int) {
	start := h1
	tail1 := h2 - 1          // 第一段尾部
	length := tail2 - h1 + 1 // 两段长度和
	tmp := []int{}
	for h1 <= tail1 || h2 <= tail2 { // 其中一段未结束
		if h1 > tail1 { // 第一段结束，处理第二段
			tmp = append(tmp, arr[h2])
			h2++
		} else if h2 > tail2 { // 第二段结束，处理第一段
			tmp = append(tmp, arr[h1])
			h1++
		} else { // 两段都未结束
			if arr[h1] <= arr[h2] {
				tmp = append(tmp, arr[h1])
				h1++
			} else {
				tmp = append(tmp, arr[h2])
				h2++
			}
		}
	}

	for i := 0; i < length; i++ { // 将排序好两段合并写入arr
		arr[start+i] = tmp[i]
	}
}
