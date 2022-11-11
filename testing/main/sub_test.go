/**
 * @author: Real
 * Date: 2022/11/12 0:49
 */
package main

import (
	"testing"
)

/**
同时测试发现累计用时比测试那两个函数用时的和要大，因为加载testing框架也要消耗时间
如果要测试单个文件，则要执行命令go test -v xx_test.go xx.go
如果需要批量执行，文件命名需要为 xxx_test.go
*/
func TestSub(t *testing.T) {
	var a int = 10
	var b int = 10
	// 减法
	subtract := Sub(a, b)
	if subtract != 5 {
		t.Fatalf("sub calculate error, correct answer is %d\n", a-b)
	}
	t.Logf("sub calculate succeed, your answer is %d\n", subtract)

	// 运行结果
	//=== RUN   TestSub
	//    sub_test.go:17: sub calculate error, correct answer is 0
	//--- FAIL: TestSub (0.00s)
	//FAIL
	//exit status 1
	//FAIL    go-study/testing/main   0.455s
}
