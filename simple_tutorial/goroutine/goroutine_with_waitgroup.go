package main

import (
	"fmt"
	"sync"
)

var helloWg sync.WaitGroup

func hello(i int) {
	// goroutine结束就登记-1
	defer helloWg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		// 启动一个goroutine就登记+1
		helloWg.Add(1)
		go hello(i)
	}
	// 等待所有登记的goroutine都结束
	helloWg.Wait()
}
