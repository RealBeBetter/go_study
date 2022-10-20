package main

import "fmt"

func main() {

	// 切片就是动态数组，是数组的一个引用，遍历，访问切片元素，获取切片长度和数组一样
	// 切片内存结构相当于一个结构体，由三部分构成：引用数组部分的首地址(ptr*)、切片长度(len)和切片容量(cap)
	// 由于是引用，所以改变切片的值，也会改变原数组的对应值
	// 切片定义的基本语法： var 变量名 [] 类型 // 例如： var a [] int
	var slice1 []int

	array := [...]int{1, 2, 3, 4, 5}
	// 根据数组创建切片
	// 定义一个切片，用切片去引用一个已经创建好的数组
	slice1 = array[1:2]
	// 修改切片值，原数组的值也会改变
	slice1[0] = 6
	slice2 := array[1:4]

	fmt.Println("数组值：", array)
	fmt.Println("切片一：", slice1)
	fmt.Println("切片二：", slice2)

	slice3 := [5]int{0, 1}
	fmt.Println("切片三：", slice3)

	// 通过make来创建切片 ： var 切片名 [] type =make([],len,[cap])
	// cap可选，要求cap >= len
	// slice0 := make([]int, 4, 10)
	var slice4 []int = make([]int, 5, 5)
	fmt.Println("切片四：", slice4)

	// 定义一个切片，直接指定具体数组，使用原理类似make的方式
	slice5 := array
	fmt.Println("切片五：", slice5)

}
