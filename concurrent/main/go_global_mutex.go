/**
 * @author: Real
 * Date: 2022/11/14 20:46
 */
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

var (
	result []int      = make([]int, 0) // 素数切片
	lock   sync.Mutex                  // 全局锁
)

func main() {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu - 1)
	// 主线程开启2000个协程，进行素数判断，等待10秒后，读取素数切片内容
	for i := 0; i < 2000; i++ {
		go isPrime(i)
	}
	// 主线程等待 10s
	time.Sleep(10 * time.Second)

	// 由于读取的是全局变量，所以读的时候也要加锁和释放锁
	for _, value := range result {
		// 请求锁
		lock.Lock()
		fmt.Println(value)
		// 释放锁
		lock.Unlock()
	}
}

func isPrime(n int) {
	isPrime := true
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime && !exist(result, n) {
		lock.Lock() // 请求锁
		result = append(result, n)
		lock.Unlock() // 释放锁
	}
}

func exist(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}
