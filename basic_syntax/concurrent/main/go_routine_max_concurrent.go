/**
 * @author: Real
 * Date: 2022/11/14 20:41
 */
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 使用内置函数获取当前机器的最大核心数
	cpu := runtime.NumCPU()
	// CPU cores:  16
	fmt.Println("CPU cores: ", cpu)
	// 设置最大线程并发数为核心数 -1
	runtime.GOMAXPROCS(cpu - 1)
}
