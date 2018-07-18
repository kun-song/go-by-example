package main

import (
	"fmt"
	"math"
)

// constant 类似 var，但声明的是常量
const s string = "constant value"

func main() {
	fmt.Println(s)

	// n 没有类型
	const n = 100000

	/**
	 * constant 可以与任意精度（arbitrary precision）的值进行算术运算
	 */
	const d = 2e20 / n
	
	fmt.Println(d)

	/**
	 * 1. constant numeric value 没有 type，除非显式给定，例如通过显式类型转换
	 */
	fmt.Println(int64(d))
	
	/**
	 * 2. constant numeric value 在特定 context 下类型可被自动推断，
	 *    例如 math.Sin 参数为 flaot64，因此此处 n 类型就是 float64
	 */
	fmt.Println(math.Sin(n))
}
