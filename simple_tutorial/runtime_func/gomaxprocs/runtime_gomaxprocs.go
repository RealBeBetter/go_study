package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	// 切换使用 1/2 个 CPU，运行结果不一样
	// runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
