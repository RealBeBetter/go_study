/**
 * @author wei.song
 * @since 2023/10/20 14:11
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello, world!")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")

	// string test
	stringTest1()
	stringTest2()

	// array test
	array()

	// slice test
	slice()

	// map test
	mapTest()

	// pointer test
	pointer()
}

func stringTest1() {
	str1 := "Golang"
	str2 := "Go语言"
	// uint8
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	// 108 l
	fmt.Println(str1[2], string(str1[2]))
	// 232 è
	fmt.Printf("%d %c\n", str2[2], str2[2])
	// len(str2)： 8
	fmt.Println("len(str2)：", len(str2))
}

func stringTest2() {
	str2 := "Go语言"
	runeArr := []rune(str2)
	// int32
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	// 35821 语
	fmt.Println(runeArr[2], string(runeArr[2]))
	// len(runeArr)： 4
	fmt.Println("len(runeArr)：", len(runeArr))
}

func array() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	// [101 102 103 104 105]
	fmt.Println(arr)
}

func slice() {
	// 创建一个 长度 为 3 的 slice
	slice1 := make([]float32, 3)
	// 创建一个长度为 3, 容量为 5 的 slice
	slice2 := make([]float32, 3, 5)
	println(len(slice1), cap(slice1))
	println(len(slice2), cap(slice2))

	// 给切片添加元素，切片容量可以根据需要自动扩展
	// [0, 0, 0, 1, 2, 3, 4]
	slice2 = append(slice2, 1, 2, 3, 4)
	println(len(slice2), cap(slice2))

	// 输出切片元素 [0 0 0 1 2 3 4]
	fmt.Println(slice2)
	// 不能输出切片元素 [7/12]0x140000161b0
	println(slice2)
	// 这里需要注意内置函数和 fmt 的打印方式的区别

	// 子切片 [start, end)
	sub1 := slice2[3:]  // [1 2 3 4]
	sub2 := slice2[:3]  // [0 0 0]
	sub3 := slice2[1:4] // [0 0 1]
	fmt.Println(sub3)
	// 合并切片
	combination := append(sub1, sub2...)
	fmt.Println(combination)

}

func mapTest() {
	// 仅声明 map, key = string, value = int
	m1 := make(map[string]int)
	// 赋值使用
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3
	fmt.Println(m1)

	// 声明时初始化
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	fmt.Println(m2)
}

func pointer() {
	str := "Golang"
	// p 是指向 str 的指针
	var p *string = &str
	// 修改了 p, str 的值也发生了改变
	*p = "Go语言"

	fmt.Println(str)
}
