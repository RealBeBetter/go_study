/**
 * @author wei.song
 * @since 2023/9/28 10:42
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abbd"))
}

func longestPalindrome(str string) string {
	if str == "" || len(str) <= 0 {
		return ""
	}

	length := len(str)
	start := 0
	maxLength := 1

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if j-i+1 > maxLength && isPalindrome(str, i, j) {
				maxLength = j - i + 1
				start = i
			}
		}
	}

	return str[start : start+maxLength]
}

func isPalindrome(str string, start int, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}
