package main

import "sort"

// https://leetcode.cn/problems/minimum-deletions-to-make-string-k-special/

// word 仅由小写字母组成
func minimumDeletions(word string, k int) int {
	// 一种策略是删除最多的，一种策略是删除最少的归零
	byteToCntMap := make(map[rune]int)
	for _, char := range word {
		byteToCntMap[char]++
	}

	counts := make([]int, 0, len(byteToCntMap))
	for _, cnt := range byteToCntMap {
		counts = append(counts, cnt)
	}
	// 排序
	sort.Ints(counts)
	if len(counts) == 1 {
		if counts[0] > k {
			return 1
		}
		return 0
	}

	// 删除字符的最小数量
	deleteCnt := 0

	// 计算 deleteCnt，寻找每一步的最优解，
	start, end := 0, len(counts)-1
	for start < end && counts[end]-counts[start] > k {
		deleteMin := counts[start]
		deleteMax := 1

		if deleteMax > deleteMin {
			// 删除第一个元素，重新定义切片，然后重新定义 start 和 end
			counts = counts[1:]
			sort.Ints(counts)
			start, end = 0, len(counts)-1

			deleteCnt += deleteMin
		} else {
			// 重新排序，然后重新定义 start 和 end
			counts[end]--
			sort.Ints(counts)
			start, end = 0, len(counts)-1

			deleteCnt += deleteMax
		}
	}

	return deleteCnt
}

func main() {
	word := "aabcaba"
	k := 0
	println(minimumDeletions(word, k))
}
