package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/**
 * Go 语言管理状态的主要机制是 communication over channels，但也提供了其他机制，比如用 `sync/atomic`
 * 管理被多个 goroutine 共享的 atomic counter
 *
 * 使用 atomic.AddUint64 和 atomic.LoadUint64 原子的更新、访问被 goroutine 共享的变量
 */

func main() {
	var counter uint64

	// 100 并发写
	for i := 0; i < 100; i++ {
		go func() {
			for {
				// 原子更新
				atomic.AddUint64(&counter, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	// 原子访问
	fmt.Println(atomic.LoadUint64(&counter))
}
