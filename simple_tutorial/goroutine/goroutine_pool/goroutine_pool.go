package main

import (
	"fmt"
	rand2 "math/rand"
)

type Job struct {
	// id
	ID int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func main() {
	// 创建任务管道和结果管道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	// 创建工作池
	createGoroutinePool(64, jobChan, resultChan)

	// 开启打印结果的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道并打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.ID,
				result.job.RandNum, result.sum)
		}
	}(resultChan)

	// 循环创建 Job，输入到任务管道
	id := 0
	for {
		id++
		randNum := rand2.Int()
		job := &Job{ID: id, RandNum: randNum}
		jobChan <- job
	}
}

// createGoroutinePool
//  @Description: 根据开的协程个数，去执行 Job 管道中的任务，输出到 Result 结果管道
//  @param goroutineNum 协程数量
//  @param jobChan
//  @param resultChan
func createGoroutinePool(goroutineNum int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < goroutineNum; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 遍历任务管道，执行运算
			for job := range jobChan {
				randNum := job.RandNum
				sum := 0
				for randNum != 0 {
					sum += randNum % 10
					randNum /= 10
				}

				// 构建结果并输出到结果管道
				result := &Result{job: job, sum: sum}
				resultChan <- result
			}
		}(jobChan, resultChan)
	}
}
