package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	/**
	 * 若 jobs 被关闭，且 jobs 中所有消息都被处理完，则 more 为 false
	 *
	 * channel 被关闭后，还能继续从中 receive 数据，不过接收到的都是 chan type 中 type 的 zero value，比如 jobs 中即为 int 的 zero value 0，
	 * 因此需要用第二个返回值，即 more 来判断 channel jobs 是否被关闭
	 */
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received a job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	// 阻塞等待
	<-done
}
