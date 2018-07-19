package main

import (
	"fmt"
)

func main() {
	x := plus(10, 20)
	fmt.Println("10 + 20 =", x)

	y := plus3(10, 20, 30)
	fmt.Println("10 + 20 + 30 =", y)
}

/**
 * 必须显式 return 返回值，不会自动以最后一个表达式的结果作为返回值（Scala、Haskell 是这样的）
 */
func plus(a int, b int) int {
	return a + b
}

/**
 * 参数连续，且类型相同时，可以只写最后一个参数的类型
 */
func plus3(a, b, c int) int {
	return a + b + c
}
