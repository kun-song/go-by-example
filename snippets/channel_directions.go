package main

import (
	"fmt"
)

/**
 * channel 用于函数参数时，可以指定消息流动的方向，更加类型安全：
 *   1. chan<- string 仅接受 send 操作
 *   2. <-chan string 仅接受 receive 操作
 */
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
