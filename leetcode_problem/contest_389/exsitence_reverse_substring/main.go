package main

import "strings"

// https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse/description/
func isSubstringPresent(s string) bool {
	builder := strings.Builder{}
	for i := range s {
		builder.WriteByte(s[len(s)-1-i])
	}

	// 判断有无两个字符相同
	reverseStr := builder.String()
	for start, end := 0, 2; end <= len(s); start, end = start+1, end+1 {
		if strings.Contains(s, reverseStr[start:end]) {
			return true
		}
	}

	return false
}
