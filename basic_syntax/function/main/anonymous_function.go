package main

import (
	"fmt"
)

// 匿名对象设置为全局变量
var (
	anonymousFunction = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	/*func sum(args...int) int{  //支持0-多个参数
		//......
	}*/

	// 2,8 表示的是传参调用
	res := func(n1 int, n2 int) int {
		return n1 * n2
	}(2, 8)
	fmt.Println("res = ", res)

	a := func(n1 int, n2 int) (int, int) {
		return n2, n1
	}
	// 匿名函数赋值给对象
	n1 := 10
	n2 := 29
	n1, n2 = a(n1, n2)

	// 调用全局匿名函数
	fmt.Println(anonymousFunction(42, 44))
}
