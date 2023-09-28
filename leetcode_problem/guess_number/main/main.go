/**
 * @author wei.song
 * @since 2023/9/28 09:25
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	result := guessNumber(10)
	fmt.Println(result)
}

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := (left + right) / 2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func guess(n int) int {
	randNumber := rand.Int()
	if randNumber > 100 {
		return -1
	} else if randNumber < 0 {
		return 1
	} else {
		return 0
	}

}
