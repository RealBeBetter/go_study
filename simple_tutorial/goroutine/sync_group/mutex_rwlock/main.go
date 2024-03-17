package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func write(useRWLock bool) {
	if useRWLock {
		rwLock.Lock() // 加写锁
		x++
		// 假设读操作耗时10毫秒
		time.Sleep(10 * time.Millisecond)
		rwLock.Unlock() // 解写锁
		wg.Done()
		return
	}

	lock.Lock() // 加互斥锁
	x++
	// 假设读操作耗时 10 毫秒
	time.Sleep(10 * time.Millisecond)
	lock.Unlock() // 解互斥锁
	wg.Done()
}

func read(useRWLock bool) {
	if useRWLock {
		rwLock.RLock()               // 加读锁
		time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
		rwLock.RUnlock()             // 解读锁
		wg.Done()
		return
	}

	lock.Lock()                  // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	useRWLock := false
	// 168.2448ms  true
	// 15.5570766s false

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(useRWLock)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read(useRWLock)
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
