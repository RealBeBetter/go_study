package main

import (
	"fmt"
)

func levenshteinDistance(s, t string) int {
	lenS := len(s)
	lenT := len(t)

	dp := make([][]int, lenS+1)
	for i := range dp {
		dp[i] = make([]int, lenT+1)
		dp[i][0] = i
	}

	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenT; j++ {
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[lenS][lenT]
}

func similarity(s, t string) float64 {
	maxLen := len(s)
	if len(t) > maxLen {
		maxLen = len(t)
	}
	distance := levenshteinDistance(s, t)
	return 1.0 - float64(distance)/float64(maxLen)
}

func main() {
	s := "kitten"
	t := "sitting"
	// 实际编辑次数为 3，则相似度为 (7-3) / 7 = 0.57
	fmt.Printf("The similarity between '%s' and '%s' is %.2f.\n", s, t, similarity(s, t))
}
