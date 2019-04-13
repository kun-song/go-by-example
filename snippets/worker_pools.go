package main

import (
	"fmt"
	"time"
)

// 与线程模型不同，worker 不是线程，也不是 Runnable 之类的，仅仅是个简单函数，并通过 goroutine 运行
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		time.Sleep(2 * time.Second)
		fmt.Println("worker", id, "finish job", j)
		result <- j * 2
	}
}

/**
 * 用 goroutine + channel 实现 worker pool
 */
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 2 个并发执行的 worker goroutine
	for i := 0; i < 2; i++ {
		go worker(i, jobs, results) // 由于没有任务，初始状态为阻塞
	}

	// 发布任务
	for i := 0; i < 8; i++ {
		jobs <- i
	}

	// 显式关闭
	close(jobs)

	for i := 0; i < 8; i++ {
		<-results
	}

	fmt.Println("finish")

	// worker 0 finish job 1
	// worker 1 finish job 0
	// worker 1 finish job 3
	// worker 0 finish job 2
	// worker 0 finish job 5
	// worker 1 finish job 4
	// worker 1 finish job 7
	// worker 0 finish job 6
	// finish
}
