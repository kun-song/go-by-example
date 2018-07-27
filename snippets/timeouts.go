package main

import (
	"fmt"
	"time"
)

/**
 * 对于需要连接外部资源，或需要指定执行时间的应用来说，超时很重要
 *
 * 超时：select + channel + time 实现
 */
func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "A"
	}()

	// 超时为 1 s
	select {
	case result := <-c1:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout: 1 second")
	}

	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "B"
	}()

	// 超时为 2s
	select {
	case result := <-c2:
		fmt.Println(result)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout: 3 seconds")
	}
}
