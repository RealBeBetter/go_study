package main

//
//import (
//	"fmt"
//)
//
//func levenshteinDistance(s, t string) int {
//	lenS := len(s)
//	lenT := len(t)
//
//	// 创建一个矩阵 dp，其中 dp[i][j] 表示 s[:i] 到 t[:j] 的最小编辑距离
//	dp := make([][]int, lenS+1)
//	for i := range dp {
//		dp[i] = make([]int, lenT+1)
//		dp[i][0] = i // 当 t 为空时，编辑距离为 i
//	}
//
//	for j := range dp[0] {
//		dp[0][j] = j // 当 s 为空时，编辑距离为 j
//	}
//
//	// 动态规划填充矩阵
//	for i := 1; i <= lenS; i++ {
//		for j := 1; j <= lenT; j++ {
//			cost := 1
//			if s[i-1] == t[j-1] {
//				cost = 0
//			}
//			dp[i][j] = min(dp[i-1][j]+1, // 删除
//				dp[i][j-1]+1,      // 插入
//				dp[i-1][j-1]+cost) // 替换
//		}
//	}
//
//	return dp[lenS][lenT]
//}
//
//func main() {
//	s := "kitten"
//	t := "sitting"
//	fmt.Printf("The Levenshtein distance between '%s' and '%s' is %d.\n", s, t, levenshteinDistance(s, t))
//}
