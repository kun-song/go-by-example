package main

import (
	"fmt"
)

func sum(xs ...int) {
	fmt.Print(xs, " ")
	sum := 0
	for x := range xs {
		sum += x
	}
	fmt.Println(sum)
}

func main() {
	// 1. 多个参数调用
	sum(1, 2, 3, 4)

	// 2. slice 调用
	xs := []int{1, 2, 3, 4}
	sum(xs...) // 有点类似 Scala
}
