/**
 * @author wei.song
 * @since 2023/10/7 11:54
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	combinations := letterCombinations("23")
	fmt.Println(combinations)
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	numberLetters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var result []string

	return backTrace(digits, 0, result, numberLetters)
}

func backTrace(digits string, index int, result []string, numberLetters []string) []string {
	if index == len(digits) {
		return result
	}

	number, _ := strconv.Atoi(digits[index : index+1])
	next := numberLetters[number]

	var nextStrings []string
	for i := 0; i < len(next); i++ {
		currentStr := string(next[i])
		if len(result) == 0 {
			nextStrings = append(nextStrings, currentStr)
		} else {
			for _, currentRes := range result {
				str := currentRes + currentStr
				nextStrings = append(nextStrings, str)
			}
		}
	}

	return backTrace(digits, index+1, nextStrings, numberLetters)
}
