package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var x int64
var l sync.Mutex

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	run(&wg, add) // 普通版add函数 不是并发安全的
	wg = sync.WaitGroup{}
	run(&wg, mutexAdd) // 加锁版add函数 是并发安全的，但是加锁性能开销大
	wg = sync.WaitGroup{}
	run(&wg, atomicAdd) // 原子操作版add函数 是并发安全，性能优于加锁版

	//9490
	//2.1559ms
	//19490
	//1.8024ms
	//29490
	//1.0466ms
}

func run(wg *sync.WaitGroup, f func()) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
