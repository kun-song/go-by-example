package main

import (
	"fmt"
)

/**
 * channel 是连接 goroutine 的通道，一个 goroutine 可以向 channel 发送消息，
 * 另一个 goroutine 可以从 channel 接收这些消息
 *
 * 默认情况下：send 和 receive 操作都是阻塞的，因为 channel 容量为 0
 */

func main() {
	// 创建 channel
	messages := make(chan string)

	// channel <- m 发送消息
	go func() {
		messages <- "ping"
	}()

	// <-channel 接收消息，阻塞，直到 receive 成功
	msg := <-messages
	fmt.Println(msg)
}
