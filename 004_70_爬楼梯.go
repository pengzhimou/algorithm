package main

// func main() {
// 	// start := time.Now() // 获取当前时间
// 	// fmt.Println(climb1(50))
// 	// elapsed := time.Since(start)
// 	// fmt.Println(elapsed)

// 	// start := time.Now()
// 	// fmt.Println(climb2(50))
// 	// elapsed := time.Since(start)
// 	// fmt.Println(elapsed)

// 	// start := time.Now()
// 	// fmt.Println(climb3(50))
// 	// elapsed := time.Since(start)
// 	// fmt.Println(elapsed)

// 	start := time.Now()
// 	fmt.Println(feb1(10))
// 	elapsed := time.Since(start)
// 	fmt.Println(elapsed, tempfeb)
// }

//    6
//  5    4
// 3 4  2 3
// 2 1 2 3 f2  1 2
// f2 f2 f2 1 2  f1 f2
//  f1 f2

// way 1
func climb1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climb1(n-1) + climb1(n-2)
}

// way 2
var temp = map[int]int{}

func climb2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if v, ok := temp[n]; ok {
		return v
	} else {
		rst := climb2(n-1) + climb2(n-2)
		temp[n] = rst
		return rst
	}
}

// way 3
func climb3(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	rst := 0
	twoStep := 2
	oneStep := 1

	for i := 3; i <= n; i++ {
		rst = twoStep + oneStep
		oneStep = twoStep
		twoStep = rst
	}
	return rst
}

// feb 1
var tempfeb = map[int]int{}

func feb1(n int) int {
	rst := 0
	if n == 0 {
		rst = 0
		tempfeb[n] = rst
		return rst
	}

	if n == 1 || n == 2 {
		rst = n
		tempfeb[n] = rst
		return rst
	}

	if v, ok := tempfeb[n]; ok {
		rst = v
	} else {
		rst = feb1(n-1) + feb1(n-2)
		tempfeb[n] = rst
	}

	return rst
}
