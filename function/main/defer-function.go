package main

import (
	"fmt"
	"strings"
)

func main() {
	f := makeSuffix(".txt")
	fmt.Println(f("szc.txt"))
	fmt.Println(f("szc"))
}

// 虽然匿名函数执行了两次，但闭包函数makeSuffix里的语句只执行了一次
// 而且defer语句先定义的后输出，且都在函数体执行完之后。
func makeSuffix(suffix string) func(string) string {
	var n = 1
	defer fmt.Println("suffix = ", suffix, ", n = ", n)
	defer fmt.Println("...")
	n = n + 1

	fmt.Println("makeSuffix..")
	return func(fileName string) string {
		if !strings.HasSuffix(fileName, suffix) {
			return fileName + suffix
		}
		return fileName
	}
}
