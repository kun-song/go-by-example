package main

import "os"

/**
 * panic 用于 fail fast，使用场景：
 *   1. 出现正常流程中绝不可能出现的错误
 *   2. 无法处理的错误
 * 出现以上错误后，没有任何合理的处理手段，只能 fail fast
 */
func main() {
	panic("a problem!")

	// 当返回值指示存在无法处理的错误时，可以用 panic
	_, err := os.Create("xx")
	if err != nil {
		panic(err)
	}
}
