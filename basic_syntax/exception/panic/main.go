package main

import "log"

func localPanic() {
	defer func() {
		log.Print("3333")
	}()
	panic("4444")
}

func main() {
	log.Print("1111")
	defer func() { log.Print("2222") }()
	localPanic()
	log.Print("9999")
}

// 2024/07/21 21:55:20 1111
// 2024/07/21 21:55:20 3333
// 2024/07/21 21:55:20 2222
// panic: 4444
//
// goroutine 1 [running]:
// main.localPanic()
//        C:/xxx/main.go:9 +0x3e
// main.main()
//        C:/xxx/main.go:15 +0x9e
//
// Process finished with the exit code 2
