package main

import (
	"fmt"
	"time"
)

/**
 * 周期任务：固定时间间隔，重复执行。
 */
func main() {
	// 类似 Timer，Ticker 的 channel 将定期被发送一个时间事件
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	fmt.Println("stop!")
}
