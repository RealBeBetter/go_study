package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// 启动 trace Goroutine
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	fmt.Println("Hello World")
}
