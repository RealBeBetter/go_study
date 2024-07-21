package main

import (
	"slices"
)

var vowels = []string{"a", "e", "i", "o", "u"}

func doesAliceWin(str string) bool {
	for {
		s, ok := oddRemove(str)
		if !ok {
			return false
		}
		if len(s) == 0 {
			return true
		}

		s, ok = evenRemove(s)
		if !ok {
			return true
		}
		if len(s) == 0 {
			return false
		}
	}
}

// oddRemove 移除奇数个元音字母，无法移除则返回原字符串
func oddRemove(s string) (string, bool) {
	maxIdx := 0
	maxCnt := 0

	for i := 0; i < len(s); i++ {
		if slices.Index(vowels, string(s[i])) != -1 {
			maxCnt++
			maxIdx = i
		}
	}

	if maxCnt == 0 {
		return s, false
	}

	if maxCnt%2 == 1 {
		return "", true
	}

	// 现存偶数个，移除
	return s[maxIdx:], true
}

// evenRemove 移除偶数个元音字母
func evenRemove(s string) (string, bool) {
	maxIdx := 0
	maxCnt := 0

	for i := 0; i < len(s); i++ {
		if slices.Index(vowels, string(s[i])) != -1 {
			maxCnt++
			maxIdx = i
		}
	}

	// 现存偶数个，全部移除
	if maxCnt%2 == 0 {
		return "", true
	}

	// 现存奇数个
	if maxIdx > 0 {
		return s[:maxIdx], true
	}
	return s[maxIdx:], false
}
