package main

import (
	"fmt"
)

/**
 * range 不但可以遍历基本数据结构，也可以用来遍历 channel 中的消息
 */
func main() {
	queue := make(chan string, 2)
	queue <- "hello"
	queue <- "world"
	// 即使 channel 被 close，但其中已有的消息还是会被处理
	close(queue)

	// range channel 时必须关闭 channel，否则 range 死循环
	for msg := range queue {
		fmt.Println(msg)
	}
}
