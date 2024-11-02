package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

// expensiveOperation 模拟一个昂贵的操作
func expensiveOperation(key string) (string, error) {
	fmt.Printf("Executing expensive operation for key: %s\n", key)
	time.Sleep(2 * time.Second) // 模拟耗时操作
	return fmt.Sprintf("Result for %s", key), nil
}

func expensiveOperationWithCtx(ctx context.Context, key string) (string, error) {
	fmt.Printf("Executing expensive operation for key: %s\n", key)
	time.Sleep(2 * time.Second) // 模拟耗时操作
	return fmt.Sprintf("Result for %s", key), nil
}

func main() {
	// 创建一个 Group 实例
	var g singleflight.Group

	// 模拟并发请求
	keys := []string{"key1", "key2", "key1"}

	for _, key := range keys {
		go func(k string) {
			// 使用 Group 的 Do 方法
			result, err, _ := g.Do(k, func() (interface{}, error) {
				// 执行昂贵的操作
				return expensiveOperation(k)
			})

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Result for %s: %s\n", k, result.(string))
			}
		}(key)
	}

	// 等待所有 goroutines 完成
	time.Sleep(3 * time.Second)

	fmt.Println("all task done.")
}

//Executing expensive operation for key: key1
//Executing expensive operation for key: key2
//Result for key2: Result for key2
//Result for key1: Result for key1
//Result for key1: Result for key1
