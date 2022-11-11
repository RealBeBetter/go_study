/**
 * @author: Real
 * Date: 2022/11/12 0:12
 */
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// 编写单元测试，单元测试的文件命名遵循规则：xxx_test.go
	sum := Add(10)
	if sum != 55 {
		t.Fatalf("calculate error , correct answer is 55, but your answer is %d\n", sum)
	}
	t.Logf("calculate succeed, your answer is %d\n", sum)

	// 使用单元测试，只需要使用 go test -v 命令即可
	//C:\Users\Real\GoProjects\go_study\testing\main>go test -v
	//=== RUN   TestAdd
	//    add_test.go:17: calculate succeed, your answer is 55
	//--- PASS: TestAdd (0.00s)
	//PASS
	//ok      go-study/testing/main   0.466s
}
