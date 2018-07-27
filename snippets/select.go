package main

/**
 * select 可以在多个 **channel 操作**（send or receive）上阻塞
 *
 * select + goroutine + channel 是 Go 中的强大组合
 */

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 阻塞 <- c1 操作 1 second
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one second"
	}()

	// 阻塞 <- c2 操作 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two second"
	}()

	/**
	 * select 在 receive 操作上阻塞（因为两个 channel 容量都是 0）
	 */
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	fmt.Println("finish")
}
