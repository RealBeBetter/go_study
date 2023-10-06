package main

import (
	"errors"
	"fmt"
)

func main() {
	// 直接报错
	//i := 5 / 0
	//fmt.Println(i)
	// 捕获异常
	test()

	// 自定义异常
	test2()
}

func test() {
	defer func() {
		// 捕获异常
		err := recover()
		if err != any(nil) {
			// 输出异常
			fmt.Println("err:", err) // err: runtime error: integer divide by zero
		}
	}()
	n1 := 1
	n2 := 0
	fmt.Println("res:", n1/n2)
}

// 自定义异常
func testError(name string) (err error) {
	if name == "szc" {
		return nil
	} else {
		return errors.New("Something wrong with " + name + "...") // 定义新的错误信息
	}
}

func test2() {
	err := testError("sss")
	if err != nil {
		panic(any(err)) // 终止程序
	}
	fmt.Println("...")
}
