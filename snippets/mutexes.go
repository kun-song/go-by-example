package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * * sync/atomic 只能用于管理简单共享状态（如计数器）
 * * mutex 用于管理复杂状态
 */
func main() {
	// 复杂共享状态：map
	state := make(map[int]int)

	// 简单共享状态：atomic counter
	var readCounter uint64
	var writeCounter uint64

	mutex := &sync.Mutex{}

	// 100 并发读
	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				k := rand.Intn(5)

				// 使用 mutex 同步共享状态读
				mutex.Lock()
				total += state[k]
				mutex.Unlock()

				// 原子写
				atomic.AddUint64(&readCounter, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 并发写
	for i := 0; i < 10; i++ {
		go func() {
			for {
				k := rand.Intn(5)
				v := rand.Intn(100)

				// 使用 mutex 同步共享状态写
				mutex.Lock()
				state[k] = v
				mutex.Unlock()

				// 原子写
				atomic.AddUint64(&writeCounter, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("readCounter = ", atomic.LoadUint64(&readCounter))
	fmt.Println("writeCounter = ", atomic.LoadUint64(&writeCounter))

	// state 正在被 100 个 goroutine 读 + 10 个 goroutine 写，此处也要同步
	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()

	// readCounter =  78784
	// writeCounter =  7880
	// map[3:8 1:45 4:7 0:87 2:93]
}
