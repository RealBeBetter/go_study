package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"testing"
	"time"
)

type testCase struct {
	name    string
	key     string
	fn      func() (interface{}, error)
	wantVal interface{}
	wantErr error
}

func TestSingleFlight(t *testing.T) {
	testCases := []testCase{
		{
			name:    "concurrent_same_key",
			key:     "key",
			fn:      func() (interface{}, error) { return "value", nil },
			wantVal: "value",
			wantErr: nil,
		},
		{
			name:    "different_key",
			key:     "key2",
			fn:      func() (interface{}, error) { return "value2", nil },
			wantVal: "value2",
			wantErr: nil,
		},
		{
			name:    "error_case",
			key:     "error_key",
			fn:      func() (interface{}, error) { return nil, fmt.Errorf("simulated error") },
			wantVal: nil,
			wantErr: fmt.Errorf("simulated error"),
		},
	}

	group := &singleflight.Group{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err, shared := group.Do(tc.key, tc.fn)
			if val != tc.wantVal || (err != tc.wantErr && err.Error() != tc.wantErr.Error()) {
				t.Errorf("For case %s, got val: %v, err: %v, want val: %v, want err: %v", tc.name, val, err, tc.wantVal, tc.wantErr)
			}
			t.Logf("For case %s, got val: %v, err: %v, shared val:%v\n", tc.name, val, err, shared)
		})
	}
}

func TestSingleFlightWithForget(t *testing.T) {
	// 创建一个 Group 实例
	var g singleflight.Group

	// 模拟并发请求
	keys := []string{"key1", "key2", "key1", "key1", "key1", "key1"}

	var wg sync.WaitGroup

	for i, key := range keys {
		wg.Add(1)
		go func(k string, index int) {
			defer wg.Done()
			if k == "key1" {
				g.Forget(k)
			}
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
		}(key, i)
	}

	wg.Wait()

	for i, key := range keys {
		go func(k string, index int) {
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
		}(key, i)
	}

	// 等待所有 goroutines 完成
	time.Sleep(3 * time.Second)

	fmt.Println("all task done.")
}

func TestSingleFlightWithDoChan(t *testing.T) {
	// 创建一个 Group 实例
	var g singleflight.Group

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	// 模拟并发请求
	keys := []string{"key1", "key2", "key1"}

	for _, key := range keys {
		go func(k string) {
			// 使用 Group 的 Do 方法
			result := g.DoChan(k, func() (interface{}, error) {
				// 执行昂贵的操作
				return expensiveOperationWithCtx(ctx, k)
			})

			select {
			case r := <-result:
				fmt.Printf("Result for %s: %s\n, error:%v", k, r.Val.(string), r.Err)
				return
			case <-ctx.Done():
				fmt.Printf("Operation timeout: %v\n", ctx.Err())
			default:
				fmt.Printf("Waiting for the result: %s\n", k)
			}
		}(key)
	}

	// 等待所有 goroutines 完成
	time.Sleep(3 * time.Second)
}
