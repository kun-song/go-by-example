package main

import (
	"fmt"
	"time"
)

/**
 * 利用 channel 可以阻塞 goroutine 的特性，可以对 goroutine 进行同步
 */

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// 利用 receive 操作，阻塞，直到 goroutine 中执行 channel <- message
	<-done
}
