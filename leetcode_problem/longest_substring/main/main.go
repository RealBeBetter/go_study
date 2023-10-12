/**
 * @author wei.song
 * @since 2023/10/12 15:31
 */
package main

import (
	"strings"
)

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbbb"))
	println(lengthOfLongestSubstring("abcdefabc"))
	println(lengthOfLongestSubstring("bbtablud"))
	println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(str string) int {
	if len(str) == 0 {
		return 0
	}

	length := len(str)
	result := 0

	// 左指针
	temp := string(str[0])
	for i := 0; i < length; i++ {
		index := strings.Index(temp, string(str[i]))

		// 存在的时候
		if index != -1 {
			temp = temp[index+1:] + string(str[i])
		} else {
			temp = temp + string(str[i])
		}
		result = maxInt(result, len(temp))
	}

	return result
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
