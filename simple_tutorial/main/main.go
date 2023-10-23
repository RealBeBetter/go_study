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
