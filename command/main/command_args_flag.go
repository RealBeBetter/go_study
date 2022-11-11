package main

import (
	"flag"
	"fmt"
)

func main() {
	// 使用 flag 包对参数进行解析
	var arg1 string
	var arg2 string
	var arg3 int
	// 接收字符串参数，参数列表：参数接收地址，参数名，默认值， 参数说明
	flag.StringVar(&arg1, "arg1", "", "第一个字符串参数")
	flag.StringVar(&arg2, "arg2", "", "第二个字符串参数")
	flag.IntVar(&arg3, "arg3", 0, "第一个整数参数")

	// 开始解析的时候必须调用
	flag.Parse()

	// 实际上是将参数的值赋值给定义好的变量
	fmt.Println("arg1 = ", arg1, ", arg2 = ", arg2, ", arg3 = ", arg3)

	// 调用的时候要指定参数名字

	// 本机运行使用结果：
	//C:\Users\Real\GoProjects\go_study\command\main>go run command_args_flag.go -arg1 字符参数1 -arg2 字符参数2 -arg3 123
	//arg1 =  字符参数1 , arg2 =  字符参数2 , arg3 =  123
}
