package main

import "fmt"

/**
 * @author wei.song
 * @since 2023/9/21 10:02
 */
func main() {
	sum := twoSum([]int{1, 2, 3, 4, 5}, 7)
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	numberMap := make(map[int]int)
	for index, num := range nums {
		tempIndex, contains := numberMap[target-num]
		if contains {
			return []int{tempIndex, index}
		}

		numberMap[num] = index
	}

	return []int{}
}
