/**
 * @author: Real
 * Date: 2022/11/12 13:20
 */
package main

import (
	"fmt"
	"sync"
)

// 使用WaitGroup等待各个协程执行完
var wg sync.WaitGroup

func main() {
	// WaitGroup 可以等待对应的协程执行完再往下执行，可以让主程序等待协程执行之后再执行
	for i := 0; i < 10; i++ {
		go show(i)
		// 每次加一个协程的数量
		wg.Add(1)
	}
	// 协程执行完毕，才会开始下面的执行任务
	wg.Wait()

	fmt.Println("sync done...")
}

func show(i int) {
	fmt.Println("show method running...", i)
	// 执行完任务线程-1
	defer wg.Done()
}
