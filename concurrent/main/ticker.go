/**
 * @author: Real
 * Date: 2022/11/14 23:56
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer只执行一次，Ticker可以周期的执行
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {
		// 每秒执行一次
		fmt.Println("ticker 1")
		counter++
		if counter >= 5 {
			break
		}
	}
	// 停止
	ticker.Stop()
}
