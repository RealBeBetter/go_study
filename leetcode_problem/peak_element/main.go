/**
 * @author wei.song
 * @since 2023/9/28 09:49
 */
package main

// 向着递增的方向找，是一定能找到的
func findPeakElement(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			// 还没达到 peak 的位置, 继续向更高的地方遍历
			start = mid + 1
		} else {
			// 右侧该值可能是 peak 的位置，所以不 +1
			end = mid
		}
	}

	return start

}
