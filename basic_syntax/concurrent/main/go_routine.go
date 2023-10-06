/**
 * @author: Real
 * Date: 2022/11/12 13:13
 */
package main

import (
	"fmt"
	"strconv"
	"time"
)

//协程从主线程开启是轻量级线程，是逻辑态的，特点：
//
//有独立的栈空间
//共享程序堆空间
//调度由用户控制
func main() {
	// 开启协程，执行指定的函数
	go testRoutine()
	for i := 0; i < 10; i++ {
		fmt.Println("main test...", strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}

	//运行结果
	//main test... 1
	//testRoutine test... 1
	//main test... 2
	//main test... 3
	//testRoutine test... 2
	//main test... 4
	//main test... 5
	//testRoutine test... 3
	//main test... 6
	//main test... 7
	//testRoutine test... 4
	//main test... 8
	//testRoutine test... 5
	//main test... 9
	//main test... 10
	//testRoutine test... 6

	// 运行结果表示，只要主线程退出，不管协程有没有完成，都会退出
}

func testRoutine() {
	for i := 0; i < 10; i++ {
		fmt.Println("testRoutine test...", strconv.Itoa(i+1))
		// 方法休眠 2s
		time.Sleep(2 * time.Second)
	}
}
