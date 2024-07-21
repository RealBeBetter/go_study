package main

import (
	"fmt"
	"strings"
)

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}

	// 使得 n == k 的最小次数
	binaryN := fmt.Sprintf("%b", n)
	binaryK := fmt.Sprintf("%b", k)

	if binaryN == binaryK {
		return 0
	}

	if len(binaryN) > len(binaryK) {
		binaryK = strings.Repeat("0", len(binaryN)-len(binaryK)) + binaryK
	} else {
		binaryN = strings.Repeat("0", len(binaryK)-len(binaryN)) + binaryN
	}

	// 只能将 n 中的 1 变为 0
	cnt := 0
	for i := 0; i < len(binaryN); i++ {
		if binaryN[i] == '1' && binaryK[i] == '0' {
			cnt++
			continue
		}

		if binaryN[i] == '0' && binaryK[i] == '1' {
			return -1
		}
	}

	return cnt
}

func main() {
	fmt.Println(minChanges(13, 4))
	fmt.Println(minChanges(21, 21))
	fmt.Println(minChanges(14, 13))
}
