package main

import (
	"fmt"
	"strings"
)

func main() {
	contains := strings.Contains("seafood", "foo")
	fmt.Println(contains)

	fmt.Println(strings.ToLower("abc"))
	println(strings.Compare("a", "b"))

	// len : 用来求长度，比如string、array、slice、map、channel
	println(len("abcdef"))
	// new : 用来分配内存，主要用来分配值类型，比如int、float32、struct…返回的是指针
	ptr := new(int)
	fmt.Println("*ptr = ", *ptr, ", ptr = ", ptr)
	*ptr = 10
	fmt.Println("*ptr = ", *ptr)
	// make: 用来分配内存，主要用来分配引用类型，比如chan、map、slice

}
