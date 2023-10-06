/**
 * @author: Real
 * Date: 2022/11/14 23:17
 */
package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// 使用 channel 写出寻找素数的方法
// 往管道里写数据时，如果超出了管道容量，就会阻塞；但是读写频率不一致，则不会发生阻塞问题。
// 不使用协程时，从空管道里读数据会发生死锁错误；
// 普通for-range遍历没有关闭的管道时，也发生死锁错误。
// 管道默认是双向的（可读可写），也可以声明为只读或者只写管道
// 使用select可以解决从管道取数据的阻塞问题
// goroutine中使用recover，解决协程中出现panic，导致程序崩溃的问题

var (
	// 待判断的数为2-80001
	initChan chan int = make(chan int, 80000)
	// 素数管道
	primeChan chan int = make(chan int, 2000)
	// 状态管道
	stateChan chan bool = make(chan bool, 4)
)

func main() {
	// 初始化数据
	go initData()
	// 获取开始时间
	start := time.Now().Unix()

	for i := 0; i < 4; i++ {
		go checkPrime()
	}

	// 匿名函数，判断是否所有判断素数的协程都已完成
	go func() {
		overNum := 0
		for {
			if overNum == 4 {
				break
			}

			status, ok := <-stateChan
			if !ok {
				break
			}

			if status {
				overNum += 1
			}
		}
		// 此时所有判断协程已经结束，关闭primeChan，主线程遍历primeChan处唤醒阻塞
		close(primeChan)
	}()

	for {
		num, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(num)
	}

	fmt.Println("Time used: ", strconv.Itoa(int(time.Now().Unix())-int(start)))
	close(stateChan)
}

// 由于有多个管道处理素数判断，所以这里最后不关闭primeChan和stateChan
func checkPrime() {
	for {
		isPrime := true
		num, ok := <-initChan
		if !ok {
			break
		}

		// 判断是否是素数
		for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeChan <- num
		}
	}

	fmt.Println("One routine has exit for the lack of data..")
	stateChan <- true
}

func initData() {
	for i := 1; i <= 80000; i++ {
		initChan <- i
	}
	// 关闭写
	close(initChan)
}
