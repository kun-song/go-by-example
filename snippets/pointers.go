package main

import (
	"fmt"
)

/**
 * Go 支持指针，但不支持指针运算
 */
func zeroval(n int) {
	n = 0
}

func zeroptr(p *int) {
	*p = 0
}

func main() {
	i := 1
	fmt.Println("initial =", i)

	zeroval(i)
	fmt.Println("zeroval:", i) // 1

	// & 获取内存地址，即指针
	zeroptr(&i)
	fmt.Println("zeroval:", i) // 0

	fmt.Println("pointer:", &i) // 0xc4200180a8
}
