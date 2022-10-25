package main

import "fmt"

func main() {
	// 追加 slice
	var slice0 []int
	slice0 = []int{0}
	fmt.Println("slice0：", slice0)

	// 直接追加元素
	slice0 = append(slice0, 6, 7, 8, 9, 10)
	fmt.Println("slice0：", slice0)

	// 直接通过数组创建
	array := [5]int{1, 2, 3, 4, 5}
	slice1 := array[0:len(array)]
	fmt.Println("slice1：", slice1)

	slice1 = append(slice1, slice0...)

	fmt.Println("slice0：", slice0)
	fmt.Println("slice1：", slice1)

	slice2 := make([]int, 5, 10)
	slice2 = append(slice2, cap(array))
	fmt.Println("slice2：", slice2)

	// 切片的拷贝，通过 copy 函数来实现
	slice3 := make([]int, 10) // 长度为10
	copy(slice3, slice2)      // 参数列表：dest、src
	fmt.Println("slice3：", slice3)
	slice2[len(slice2)-1] = -1
	// 修改源 slice 的值，拷贝之后的新 slice 的值不变
	// slice 的拷贝是一个深拷贝
	fmt.Println(slice3) // 原切片不变

	str := "hello@world!"
	slice := str[:]
	fmt.Println("slice=", slice)
	// string 是不可变的，不能通过str[0] ='z' 方式来修改字符串
	// 修改字符串，可以将 string -> []byte  或者 []rune  -> 修改  ->重写转成string
	upStr := []byte(str)

	// 可以处理数字和英文还有符号，不能处理中文。原因：一个汉字占3个字节
	upStr[5] = '-'
	str = string(upStr)
	fmt.Println("str=", str)

	// 处理中文 转成 []rune 即可
	chStr := []rune(str)
	chStr[5] = '一'
	chStr = []rune(string(chStr))
	fmt.Println("chStr=", chStr)

}
