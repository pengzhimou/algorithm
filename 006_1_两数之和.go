package main

import "fmt"

// func main() {
// 	// twoAdd1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
// 	twoAdd2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
// }

func twoAdd1(lst []int, target int) {
	temp := map[int]int{}
	for idx, i := range lst {

		temp[i] = idx
		for it, idxt := range temp {
			if idx == idxt {
				continue
			}
			if i+it == target {
				delete(temp, it)
				fmt.Printf("Num: %v %v, Index: %v %v\n", i, it, idx, idxt)
			}
		}
	}
}

func twoAdd2(lst []int, target int) []int {
	temp := map[int]int{}
	for idx, i := range lst {

		if idxt, ok := temp[target-i]; ok {
			return []int{idxt, idx}
		}
		temp[i] = idx
	}
	return nil
}
