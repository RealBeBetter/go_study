package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int64
	var wg sync.WaitGroup

	wg.Add(2)
	go add(&x, &wg)
	go add(&x, &wg)
	wg.Wait()

	fmt.Println(x)
}

func add(x *int64, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*x++
	}
	wg.Done()
}
