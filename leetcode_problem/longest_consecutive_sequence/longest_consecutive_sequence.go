/**
 * @author Real
 * @since 2023/12/19 21:53
 */
package main

// longestConsecutive
//  @Description: https://leetcode.cn/problems/longest-consecutive-sequence
//  @param nums
//  @return int
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0
	for number := range numSet {
		// 防止连续的序列需要重新计算
		if !numSet[number-1] {
			currentNumber := number
			currentStreak := 1
			for numSet[currentNumber+1] {
				currentNumber++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
