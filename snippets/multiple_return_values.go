package main

import (
	"fmt"
)

/**
 * Go 内置支持多个返回值，这是地道的 Go 用法，例如同时返回结果和错误信息
 */
func main() {
	x, y := vals()
	fmt.Println(x, y)

	_, c := vals()
	fmt.Println(c)
}

func vals() (int, int) {
	return 10, 20
}
