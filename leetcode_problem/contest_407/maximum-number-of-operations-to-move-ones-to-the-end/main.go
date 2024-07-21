package main

import (
	"fmt"
)

func maxOperations(s string) int {
	count := 0
	ans := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			count++
		}
		if s[i] == '1' && s[i+1] == '0' {
			ans += count
		}
	}

	return ans
}

func main() {
	s := "0010000111"
	fmt.Println(maxOperations(s))
}
