package main

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	var cnt int64
	products := make([]int, len(nums2))
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if i == 0 {
				products[j] = nums2[j] * k
			}

			if nums1[i]%(products[j]) == 0 {
				cnt++
			}
		}
	}
	return cnt
}

func numberOfPairs2(nums1 []int, nums2 []int, k int) int64 {
	ans := 0
	cnt1, cnt2 := map[int]int{}, map[int]int{}
	for _, num := range nums1 {
		cnt1[num]++
	}
	for i := range nums2 {
		cnt2[nums2[i]*k]++
	}
	for num, cnt := range cnt2 {
		t := num
		for t <= 1_000_000 {
			ans += cnt * cnt1[t]
			t += num
		}
	}
	return int64(ans)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{1, 2, 3, 4, 5, 6}
	k := 2
	fmt.Println(numberOfPairs2(nums1, nums2, k))
}
