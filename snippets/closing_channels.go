package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	/**
	 * 若 jobs 被关闭，且 jobs 中所有消息都被处理完，则 more 为 false
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
