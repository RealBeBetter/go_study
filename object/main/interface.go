package main

import "fmt"

func main() {

	//type 接口名 interface{
	//	method1(参数列表) 返回值列表
	//	method2(参数列表) 返回值列表
	//}
}

type ICalculate interface {
	add(a int, b int) int
	sub(a int, b int) int
}

// C 定义两个结构体，实现该接口
type C struct {
}

type D struct {
}

func (c C) add(a int, b int) int {
	fmt.Println("C...add")
	return a + b
}
