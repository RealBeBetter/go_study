/**
 * @author Real
 * @since 2023/10/28 15:57
 */
package main

import "testing"

// 运行 go test，将自动运行当前 package 下的所有测试用例，如果需要查看详细的信息，可以添加 -v 参数
func TestAdd(t *testing.T) {
	if ans := calculateAdd(1, 2); ans != 3 {
		t.Errorf("calculateAdd(1, 2) should be 3, but actually is %d", ans)
	}
}

//PS C:\Users\Real\GoProjects\go_study\simple_tutorial\testing> go test -v
//=== RUN   TestAdd
//--- PASS: TestAdd (0.00s)
//PASS
//ok      go-study/simple_tutorial/testing        0.724s
