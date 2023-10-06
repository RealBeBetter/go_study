package main

import "fmt"

func main() {

	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(3)) // 14
	fmt.Println(f(3)) // 17

}

// AddUpper 由于形成了匿名函数+外部引用的形式，所以每次调用AddUpper()时，n都会继承上一次调用的值。
func AddUpper() func(int) int {
	// 可以看做n是AddUpper()的属性，一个对象只会初始化一次。
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}
