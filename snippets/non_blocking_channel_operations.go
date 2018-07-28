package main

import (
	"fmt"
)

/**
 * channel send 和 receive 操作默认都是阻塞（blocking）的，通过 select + default 子句可实现非阻塞
 * 的 send/receive 操作
 */
func main() {
	messages := make(chan string)
	signals := make(chan string)

	// 非阻塞的 receive 操作
	select {
	case msg := <-messages:
		fmt.Println("received message = ", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞的 send 操作
	msg := "Hello"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 非阻塞的多路 channel 操作
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
