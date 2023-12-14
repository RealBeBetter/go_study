/**
 * @author wei.song
 * @since 2023/10/6 10:28
 */
package main

import (
	"strconv"
)

func main() {
	println(isPalindrome(-121))
	println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	str := strconv.Itoa(x)

	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}
