package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启 Goroutine 将 0~100 的数发送到 ch1 中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启 Goroutine 从 ch1 中接收值，并将该值的平方发送到 ch2 中
	go func() {
		for {
			// 通道关闭后再取值时 ok == false
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主 Goroutine 中从 ch2 中接收值打印
	// 通道关闭后会退出 for-range 循环
	for i := range ch2 {
		fmt.Println(i)
	}
}
