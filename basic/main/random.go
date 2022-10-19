package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //time.Now().UnixNano() 返回当前系统的纳秒数
	n := rand.Intn(100) + 1
	fmt.Println("n=", n)
}
