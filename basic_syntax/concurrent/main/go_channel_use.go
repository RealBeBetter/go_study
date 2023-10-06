/**
 * @author: Real
 * Date: 2022/11/14 23:44
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	readOnlyChan  chan<- int = make(chan int, 8) // 只写
	writeOnlyChan <-chan int = make(chan int, 8) // 只读

	intChan chan int    = make(chan int, 10)
	strChan chan string = make(chan string, 10)
)

func main() {
	// go中的channel在使用传统方法遍历管道时，如果管道不关闭，就会发生死锁
	go initIntChan()
	go initStrChan()

	wg := sync.WaitGroup{}

	wg.Wait()
	// 如果我们不确定何时关闭管道，就可以使用select来进行判断
	// 并且搭配标识符，进行循环跳出
label:
	for {
		select {
		case v := <-intChan:
			// 如果管道一直不关闭，也不会死锁，而会向下匹配
			fmt.Println("data from int chan: ", v)
		case v := <-strChan:
			fmt.Println("data from string chan: ", v)
		default:
			break label
		}
	}
}

func initIntChan() {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
}

func initStrChan() {
	for i := 0; i < 5; i++ {
		strChan <- "string " + strconv.Itoa(i)
	}
}
