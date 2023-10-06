/**
 * @author: Real
 * Date: 2022/11/14 23:54
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer顾名思义，就是定时器的意思，可以实现一些定时操作
	// 内部也是通过channel来实现的。Timer只执行一次
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)

	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	// 如果只是想单纯的等待的话，可以使用 time.Sleep 来实现
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("2s后")

	time.Sleep(time.Second * 2)
	fmt.Println("再一次2s后")

	// time.After函数的返回值是chan Time
	<-time.After(time.Second * 2)
	fmt.Println("再再一次2s后")

	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer 3 expired")
	}()

	// 停止定时器
	stop := timer3.Stop()
	// 阻止timer事件发生，当该函数执行后，timer计时器停止，相应的事件不再执行
	if stop {
		fmt.Println("Timer 3 stopped")
	}

	fmt.Println("before")
	// 原来设置5s
	timer4 := time.NewTimer(time.Second * 5)
	// 重新设置时间,即修改NewTimer的时间
	timer4.Reset(time.Second * 1)
	<-timer4.C
	fmt.Println("after")
}
