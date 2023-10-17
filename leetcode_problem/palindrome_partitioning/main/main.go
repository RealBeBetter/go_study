/**
 * @author wei.song
 * @since 2023/10/17 12:08
 */
package main

func partition(s string) [][]string {
	if s == "" || len(s) == 0 {
		return [][]string{}
	}

	result := make([][]string, 0)
	backTracking(&result, []string{}, s, 0)
	return result
}

func backTracking(result *[][]string, path []string, s string, startIndex int) {
	if len(s) == startIndex {
		*result = append(*result, append([]string{}, path...))
		return
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
			backTracking(result, path, s, i+1)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
