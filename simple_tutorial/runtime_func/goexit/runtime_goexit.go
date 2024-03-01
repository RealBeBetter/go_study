package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程 从这里往下的部分都不会执行
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	time.Sleep(1 * time.Second)
}
