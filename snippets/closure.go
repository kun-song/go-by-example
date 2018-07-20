package main

import (
	"fmt"
)

/**
 * Go 支持匿名函数，进而支持闭包
 */
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	nextInt = intSeq()
	fmt.Println(nextInt()) // 1
}
