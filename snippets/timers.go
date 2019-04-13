package main

import (
	"fmt"
	"time"
)

/**
 * 定时任务：在未来某时刻执行代码，Go 用 Timer 实现定时任务。
 */
func main() {
	// timer1 代表未来某时刻的 **单个事件**，2s 后，会向 timer1.C channel 发送一个事件
	timer1 := time.NewTimer(2 * time.Second)

	// 阻塞在 Timer 的 Channel 上
	<-timer1.C
	fmt.Println("2 seconds passed")

	// 若只想要阻塞作用，time.Sleep 更加合适，但 timer 可以取消
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("1 second passed")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 has benn stopped")
	}

	// 2 seconds passed
	// timer2 has benn stopped
}
