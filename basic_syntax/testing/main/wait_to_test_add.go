/**
 * @author: Real
 * Date: 2022/11/12 0:05
 */
package main

func Add(n int) int {
	// 等待测试文件，编写单元测试的前提
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
