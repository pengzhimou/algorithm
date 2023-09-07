package main

import "fmt"

func main() {
	// twoAdd1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	twoAdd2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func twoAdd1(lst []int) {
	temp := map[int]int{}
	for idx, i := range lst {

		temp[i] = idx
		for it, idxt := range temp {
			if idx == idxt {
				continue
			}
			if i+it == 10 {
				delete(temp, it)
				fmt.Printf("Num: %v %v, Index: %v %v\n", i, it, idx, idxt)
			}
		}
	}
}

func twoAdd2(lst []int) {
	temp := map[int]int{}
	for idx, i := range lst {

		temp[i] = idx

	}
}
