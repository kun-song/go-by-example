package main

import (
	"fmt"
)

/**
 * 默认情况下，channel 无缓存，即 send（channel <-）操作会一直阻塞，直到
 * receive（<- channel）操作成功
 *
 * 而 buffered channel 可以根据容量大小，在无 receiver 的情况下，也可 send 成功（不阻塞）
 */
func main() {
	messages := make(chan string, 2)

	messages <- "x"
	messages <- "y"
	// 因此 buffer 容量为 2，所以如果有第三个 send，则会阻塞
	// messages <- "z"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
