/**
 * @author wei.song
 * @since 2023/10/6 15:13
 */
package main

import "strings"

func main() {
	strs := []string{"flower", "flow", "flight"}
	println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	minLength := len(strs[0])

	length := len(strs)
	for i := 0; i < length; i++ {
		currentLength := len(strs[i])
		if currentLength < minLength {
			minLength = currentLength
		}
	}

	if minLength == 0 {
		return ""
	}

	builder := strings.Builder{}

	for i := 0; i < minLength; i++ {
		char := strs[0][i]
		for j := 0; j < length; j++ {
			if char != strs[j][i] {
				return builder.String()
			}
		}
		builder.WriteByte(char)
	}

	return builder.String()
}
