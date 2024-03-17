package main

// https://leetcode.cn/problems/count-substrings-starting-and-ending-with-given-character/description/

// 双指针，超时
func countSubstringsWithDoublePointer(s string, c byte) int64 {
	count := int64(0)

	for i := range s {
		if s[i] != c {
			continue
		}

		start, end := i, i
		for start >= 0 && end < len(s) {
			if s[start] == s[end] && end >= start {
				count++
			}

			end++
		}
	}

	return count
}

// 组合算法，通过
func countSubstrings(s string, c byte) int64 {
	cnt := 0

	for i := range s {
		if s[i] != c {
			continue
		}
		cnt++
	}

	if cnt < 2 {
		return int64(cnt)
	}

	return int64(cnt*(cnt-1)/2) + int64(cnt)
}
