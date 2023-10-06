/**
 * @author: Real
 * Date: 2022/11/14 23:59
 */
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 当我们需要在某个协程函数里捕获异常时，使用以前的defer-recover即可
	// 开启协程
	go testException()
	go test()
	//
	time.Sleep(time.Second * 10)
}

func testException() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test catch an exception:", err)
		}
	}()
	var map0 map[int]string
	map0[0] = "Real"
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test test.......", strconv.Itoa(i+1))
		time.Sleep(500 * time.Millisecond)
	}
}
