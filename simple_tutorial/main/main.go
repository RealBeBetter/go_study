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
}
