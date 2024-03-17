package main

import (
	"fmt"
	"time"
)

func main() {
	// 判断 channel 是否满
	strChan := make(chan string, 10)
	go write(strChan)

	for str := range strChan {
		fmt.Println("res:", str)
		time.Sleep(time.Second)
	}

	//write hello
	//res: hello
	//write hello
	//channel is full
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			println("write hello")
		default:
			println("channel is full")
		}

		time.Sleep(500 * time.Millisecond)
	}
}
