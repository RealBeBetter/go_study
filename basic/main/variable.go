package main

import "fmt"

// 函数外声明全局变量
var (
	n1 = 1
	n2 = 2
	n3 = 3
)

var n4 = "n4"

func main() {
	// var 变量名 类型 = 值
	var i int = 10
	// var 变量名 = 值
	var j = 1.2
	// 变量名 := 值，自动推导类型
	name := "szc"

	fmt.Println("i = ", i, ", j = ", j, ", name = ", name)

	// 一次声明多个变量，变量名和值一一对应
	var a, sex, b = 1, "male", 7
	// 变量声明后必须使用，而且不能隐式改变类型(int转float)
	fmt.Println("a = ", a, ", sex = ", sex, ", b = ", b)
	// a, sex, b := 2, "male", 4
	// fmt.Println("a = ", a, ", sex = ", sex, ", b = ", b)

	fmt.Println("n1 = ", n1, ", n2 = ", n2, "n3 = ", n3, ", n4 = ", n4)

}
