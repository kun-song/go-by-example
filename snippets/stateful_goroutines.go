package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/**
 * 使用锁（mutex）当然可以对共享变量的访问进行同步，但 Go 的 channle/goroutine 其实内置了同步语义，而且 Go
 * 推崇 share memory by communicating and have each piece of data owned by exactly 1 goroutine，
 *
 * 即：不使用显式锁，而是通过 goroutine 之间通过 channel 通信实现 **共享内存**
 *
 * 下面例子中的 channel + goroutine 实现方式比 mutex 烦琐一些，但当 mutex or channel 太多时，
 * 保证 mutex 使用的正确性会越来越困难，因此两种方案视实际情况选择最合适的即可，可读性更重要
 */

type readOp struct {
	k    int
	resp chan int
}

type readAllOp struct {
	resp chan map[int]int
}

type writeOp struct {
	k    int
	v    int
	resp chan bool
}

func main() {
	var readCounter uint64
	var writeCounter uint64

	// 下面 3 类 goroutine 之间，通过共享 channel 通信
	reads := make(chan *readOp)
	readAlls := make(chan *readAllOp)
	writes := make(chan *writeOp)

	// 对 map 的查询、更新都在一个 goroutine 中实现
	go func() {
		// 状态是 goroutine 私有的，不被共享，所以无需 mutex 同步
		state := make(map[int]int)

		// 通过死循环，不断接收请求
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.k]
			case readAll := <-readAlls:
				readAll.resp <- state
			case write := <-writes:
				state[write.k] = write.v
				write.resp <- true
			}
		}
	}()

	// 100 并发读
	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := &readOp{rand.Intn(5), make(chan int)}
				reads <- read

				// 因为 read.resp 容量为 0，所以上面的 goroutine 写入响应是阻塞，直到下面显式 receive 后恢复
				<-read.resp

				atomic.AddUint64(&readCounter, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 并发写
	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := &writeOp{rand.Intn(5), rand.Intn(100), make(chan bool)}
				writes <- write

				// 因为 write.resp 容量为 0，所以上面的 goroutine 写入响应是阻塞，直到下面显式 receive 后恢复
				<-write.resp

				atomic.AddUint64(&writeCounter, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readAll := &readAllOp{make(chan map[int]int)}
	readAlls <- readAll
	state := <-readAll.resp

	fmt.Println(state)

	fmt.Println(atomic.LoadUint64(&readCounter))
	fmt.Println(atomic.LoadUint64(&writeCounter))
}
