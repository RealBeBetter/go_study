package main

import "log"

func fatal() {
	defer func() {
		log.Print("3333")
	}()
	log.Fatal("4444")
}

func main() {
	log.Print("1111")
	defer func() { log.Print("2222") }()
	fatal()
	log.Print("9999")
}

// 2024/07/21 21:54:12 1111
// 2024/07/21 21:54:12 4444
//
// Process finished with the exit code 1
