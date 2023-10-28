/**
 * @author Real
 * @since 2023/10/28 15:34
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("-----------goroutine test-----------")
	downloadWithWaitGroup()

	fmt.Println("-----------goroutine test-----------")
	downloadWaitChannelMsg()
}

func download(url string) {
	fmt.Println("start to download", url)
	// 模拟耗时操作
	time.Sleep(time.Second)
	// 下载完毕则减去一个计数
	fmt.Println("download finished", url)
	wg.Done()
}

// wait group is used to wait for all goroutine finished
// 类似于 Java 中的 CountDownLatch
func downloadWithWaitGroup() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		// go is a keyword to create a goroutine
		go download("a.com/" + string(rune(i+'0')))
	}
	// wait until all goroutine finished
	wg.Wait()
	fmt.Println("all download finished!")
}

// channel is used to communicate between goroutine
var ch = make(chan string, 10)

func downloadWithChannel(url string) {
	fmt.Println("start to download", url)
	// 模拟耗时操作
	time.Sleep(time.Second)
	// send message to channel
	fmt.Println("download finished", url)
	ch <- url
}

func downloadWaitChannelMsg() {
	for i := 0; i < 3; i++ {
		go downloadWithChannel("a.com/" + string(rune(i+'0')))
	}

	// receive message from channel
	for i := 0; i < 3; i++ {
		// 使用 channel 时，一定要保证 goroutine 先执行，否则会出现 deadlock
		// 也就是说，channel 的发送和接收必须在不同的 goroutine 中进行
		// 如果在同一个 goroutine 中，那么就会出现 deadlock
		channelMsg := <-ch
		// 使用 channel 信道，可以在协程之间传递消息。阻塞等待并发协程返回消息
		fmt.Println("receive message from channel:", channelMsg)
	}
	fmt.Println("all download finished!")
}
