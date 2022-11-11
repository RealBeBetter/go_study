/**
 * @author: Real
 * Date: 2022/11/12 0:12
 */
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// 编写单元测试
	sum := Add(10)
	if sum != 55 {
		t.Fatalf("calculate error , correct answer is 55, but your answer is %d\n", sum)
	}
	t.Logf("calculate succeed, your answer is %d\n", sum)

	// 使用单元测试，只需要使用 go test -v 命令即可
}
