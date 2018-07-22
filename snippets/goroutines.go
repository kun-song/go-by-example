package main

import (
	"fmt"
)

/**
 * goroutine is a lightweight thread of execution
 */

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 直接调用方法
	f("direct")

	// go 关键字，创建 goroutine 运行 f 方法
	go f("goroutine")

	// goroutine 运行匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// This Scanln requires we press a key before the program exits.
	fmt.Scanln()
	fmt.Println("done")
}
