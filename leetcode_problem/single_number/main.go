/**
 * @author wei.song
 * @since 2023/10/16 17:29
 */
package main

func singleNumber(nums []int) []int {
	resMap := map[int]int{}
	for _, num := range nums {
		value, contains := resMap[num]
		if contains {
			value++
			resMap[num] = value
		} else {
			resMap[num] = 1
		}
	}

	var res []int
	for key, value := range resMap {
		if value == 1 {
			res = append(res, key)
		}
	}

	return res
}
