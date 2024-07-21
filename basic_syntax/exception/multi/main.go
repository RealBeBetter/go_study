package main

import (
	"log"
	"time"
)

func fatal() {
	defer func() {
		log.Print("3333")
	}()
	panic("4444")
}

func main() {
	log.Print("1111")
	defer func() { log.Print("2222") }()
	// panic 会在子 Goroutine 向上传递到主 Goroutine
	go fatal()
	time.Sleep(1 * time.Second)
	log.Print("9999")
}

// 2024/07/21 23:04:15 1111
// 2024/07/21 23:04:15 3333
// panic: 4444
//
// goroutine 5 [running]:
// main.fatal()
//        C:/xxx/main.go:12 +0x3e
// created by main.main in goroutine 1
//        C:/xxx/main.go:18 +0xa5
//
// Process finished with the exit code 2
