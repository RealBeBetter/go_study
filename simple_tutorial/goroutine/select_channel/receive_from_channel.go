package main

import "time"

func test1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 4)
	ch <- "test2"
}

func main() {
	outputChan1 := make(chan string)
	outputChan2 := make(chan string)

	// 开启 2 个 goroutine，每个 goroutine 从不同的 channel 中接收数据
	go test1(outputChan1)
	go test2(outputChan2)

	// 使用 select 语句从 2 个 channel 中接收数据
	select {
	case msg1 := <-outputChan1:
		println(msg1)
	case msg2 := <-outputChan2:
		println(msg2)
	}
}
