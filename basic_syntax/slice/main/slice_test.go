package main

import "fmt"

func main() {
	var array = [...]int{24, 69, 80, 57, 13, 34, 1, 9, 56, 90, 44, 11, 16, 78, 99, 4}
	fmt.Println("排序前：", array)
	bubbleSort(array[:])
	fmt.Println("冒泡排序后：", array)
	var target = 5
	index := binarySearch(target, array[:], 0, len(array))
	fmt.Println("二分查找 target =", target, "，数组 array =", array)
	fmt.Println("二分查找结果下标为： index =", index)
}

func bubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				var a = array[j]
				array[j] = array[j+1]
				array[j+1] = a
			}
		}
	}
}

func binarySearch(target int, array []int, left int, right int) int {
	if left > right {
		return -1
	}
	var mid = left + (right-left)/2
	if array[mid] == target {
		fmt.Println("二分查找成功，结果：", mid)
		return mid
	} else if array[mid] < target {
		binarySearch(target, array, mid+1, right)
	} else {
		binarySearch(target, array, left, mid-1)
	}
	return -1
}
