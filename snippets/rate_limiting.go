package main

import (
	"fmt"
	"time"
)

/**
 * 限流（rate limiting）是一种控制资源利用率、维持服务质量的重要机制
 *
 * Go 通过 goroutine + channel + ticker 优雅实现限流
 */
func main() {
	/**
	 * 1. 简单限流，每 200ms 处理一个请求
	 */

	// 任务
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// 限流器，类型为 chan time.Time，每 200ms 收到一个 time.Time 事件
	limiter := time.Tick(200 * time.Millisecond)

	// 使用
	for v := range requests {
		// 通过在 <-limiter 上阻塞，实现限流
		<-limiter
		fmt.Println(v, time.Now())
	}

	// 0 2019-04-13 09:00:16.470550285 +0800 CST m=+0.201347119
	// 1 2019-04-13 09:00:16.671667845 +0800 CST m=+0.402463177
	// 2 2019-04-13 09:00:16.8734545 +0800 CST m=+0.604248325
	// 3 2019-04-13 09:00:17.070800584 +0800 CST m=+0.801592935
	// 4 2019-04-13 09:00:17.274784445 +0800 CST m=+1.005575273

	/**
	 * 2. 复杂限流，可以处理 3 个洪峰
	 */

	// 任务
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 限流器
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// burstyLimiter channel 容量为 3，所以下面的 <- 阻塞，直到有空位可以插入
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 使用
	for v := range burstyRequests {
		<-burstyLimiter
		fmt.Println(v, time.Now())
	}

	// 0 2019-04-13 09:00:17.275076179 +0800 CST m=+1.005867005
	// 1 2019-04-13 09:00:17.275125124 +0800 CST m=+1.005915950
	// 2 2019-04-13 09:00:17.275149666 +0800 CST m=+1.005940491
	//
	// 3 2019-04-13 09:00:17.476337621 +0800 CST m=+1.207126944
	// 4 2019-04-13 09:00:17.676279714 +0800 CST m=+1.407067544
}
