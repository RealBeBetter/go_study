package main

import (
	"fmt"
	"os"
)

func main() {
	// 命令行参数

	// 命令行参数可以直接使用 os.Args 进行读取
	args := os.Args
	for index, arg := range args {
		fmt.Println("第 ", index+1, " 个参数是 ", arg)
	}

	// 运行的时候需要需要使用特定的命令行来完成调用
	// go run command_args.go xxx xxx

	// 本机使用命令
	//C:\Users\Real\GoProjects\go_study\command\main>go run command_args.go firstArg secondArg
	//第  1  个参数是  C:\Users\Real\AppData\Local\Temp\go-build090102442\b001\exe\command_args.exe
	//第  2  个参数是  firstArg
	//第  3  个参数是  secondArg
}
