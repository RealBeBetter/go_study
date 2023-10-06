/**
 * @author: Real
 * Date: 2022/11/14 21:08
 */
package main

import (
	"fmt"
	"sync"
)

var (
	writeChan chan int  = make(chan int, 50) // 数据管道，整型管道，容量50(不可扩容)
	exitChan  chan bool = make(chan bool, 1) // 状态管道，布尔型管道，容量1(不可扩容)
	locker    sync.Mutex
)

func main() {
	// 管道 channel 为引用类型，必须先初始化才能使用；本质是一个队列，有类型，而且线程安全
	locker.Lock()
	go writeData()
	locker.Unlock()
	go readData()

	for {
		value, ok := <-exitChan
		if !ok && value == true {
			break
		}

	}
}

// 构造读写函数。写函数往数据管道里写入50个数据，并关闭数据管道；
func writeData() {
	for i := 0; i < 50; i++ {
		// 倒序指针符号，表示向管道中插入数据
		writeChan <- i
		fmt.Println("write data: ", i)
	}
	// 关闭管道，管道不可再写
	close(writeChan)
}

// 读函数负责从数据管道里读取数据，如果读完，则往状态管道里置入true，表示读取完成
func readData() {
	for {
		value, ok := <-writeChan
		if !ok {
			// 如果管道为空，则 OK 为 false 值
			break
		}
		fmt.Println("read data: ", value)
	}
	exitChan <- true
	close(exitChan)
}
