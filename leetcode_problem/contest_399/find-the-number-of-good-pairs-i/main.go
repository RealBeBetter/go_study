package main

import "fmt"

// numberOfPairs
//  @Description: 优质数对的总数，如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对
//  @param nums1
//  @param nums2
//  @param k
//  @return int
func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	cnt := 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{1, 3, 4}
	k := 1

	pairs := numberOfPairs(nums1, nums2, k)
	fmt.Println(pairs)
}
