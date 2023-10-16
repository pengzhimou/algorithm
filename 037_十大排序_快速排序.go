package main

func main() {

}

func QuickSort1(arr []int) []int {
	return _QuickSort1(arr, 0, len(arr)-1)
}

func _QuickSort1(arr []int, left int, right int) []int {
	if left < right {
		partitionIndex := Partition1Way(arr, left, right)
		// partitionIndex := Partition2Way(arr, left, right)
		_QuickSort1(arr, left, partitionIndex-1)
		_QuickSort1(arr, partitionIndex+1, right)
	}
	return arr
}

// 快速排序--单路
func Partition1Way(arr []int, left int, right int) int {
	// 先分区，最后把基准换到边界上
	privot := left
	index := privot + 1
	for i := index; i <= right; i++ {
		if arr[privot] > arr[i] { // 当前值小于基准就交换，大于的不用管
			arr[index], arr[i] = arr[i], arr[index]
			index++ // 交换后的下一个
		}
	}
	// arr[index]是大于基准的
	arr[privot], arr[index-1] = arr[index-1], arr[privot]
	return index - 1
}

//////////////////////////////////////////

func QuickSort2(arr []int) []int {
	return _QuickSort1(arr, 0, len(arr)-1)
}

func _QuickSort2(arr []int, left int, right int) []int {
	if left < right {
		partitionIndex := Partition2Way(arr, left, right)
		// partitionIndex := Partition2Way(arr, left, right)
		_QuickSort2(arr, left, partitionIndex-1)
		_QuickSort2(arr, partitionIndex+1, right)
	}
	return arr
}

// 快速排序--双路版
func Partition2Way(arr []int, low int, high int) int {
	tmp := arr[low] // 基准
	for low < high {
		// 当队尾的元素大于等于基准数据时,向前挪动high指针
		for low < high && arr[high] >= tmp {
			high--
		}
		// 如果队尾元素小于tmp了,需要将其赋值给low
		arr[low] = arr[high]
		// 当队首元素小于等于tmp时,向前挪动low指针
		for low < high && arr[low] <= tmp {
			low++
		}
		// 当队首元素大于tmp时,需要将其赋值给high
		arr[high] = arr[low]

	}
	// 跳出循环时low和high相等,此时的low或high就是tmp的正确索引位置
	// 由原理部分可以很清楚的知道low位置的值并不是tmp,所以需要将tmp赋值给arr[low]
	arr[low] = tmp
	return low
}

//////////////////////////////////////////

// 快速排序--三路
func QuickSort3Way(arr []int) []int {
	// 确定分区位置
	return _QuickSort3Way(arr, 0, len(arr)-1)
}
func _QuickSort3Way(arr []int, left int, right int) []int {
	if left < right {
		lo, gt := Partition3Way(arr, left, right)
		_QuickSort3Way(arr, left, lo-1)
		_QuickSort3Way(arr, gt, right)
	}
	return arr
}
func Partition3Way(arr []int, left, right int) (int, int) {
	key := arr[left]
	lo, gt, cur := left, right+1, left+1 // lo和gt是相等区左右边界
	for cur < gt {
		if arr[cur] < key { // 小于key，移到前面
			arr[cur], arr[lo+1] = arr[lo+1], arr[cur] // lo+1,保证最后arr[lo]小于key
			lo++                                      // 左边界右移
			cur++                                     // 能够确定换完之后该位置值小于key,
		} else if arr[cur] > key {
			arr[cur], arr[gt-1] = arr[gt-1], arr[cur]
			gt-- // 从后面换到前面，不知道是否比key的大，还要再比一下，所以cur不移动
		} else {
			cur++
		}
	}
	arr[left], arr[lo] = arr[lo], arr[left] // 最后移动基准，arr[lo]一定比key小
	return lo, gt
}
