/**
 * @author Real
 * @since 2023/12/31 10:40
 */
package main

import (
	"strconv"
	"strings"
)

//
// hasTrailingZeros https://leetcode.cn/problems/check-if-bitwise-or-has-trailing-zeros/description/
//
//  @Description: 检查按位与是否存在尾随 0
//  @param nums
//  @return bool
//
func hasTrailingZeros(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			number := nums[i] | nums[j]
			binaryNumber := strconv.FormatInt(int64(number), 2)
			if strings.LastIndex(binaryNumber, "0") == len(binaryNumber)-1 {
				return true
			}
		}
	}

	return false
}
