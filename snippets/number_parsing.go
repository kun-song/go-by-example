package main

import (
	"fmt"
	"strconv"
)

/**
 * `strconv` 包：string -> number
 */
func main() {
	p := fmt.Println

	f, _ := strconv.ParseFloat("1.234", 64) // 精度：64
	p(f)                                    // 1.234‘

	// 0：自定推断进制
	i, _ := strconv.ParseInt("123", 0, 64)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(i, d) // 123 456

	// Atoi 进行 10 进制数字字符串转换的方便函数
	b, _ := strconv.Atoi("123")
	p(b) // 123

	_, e := strconv.Atoi("hello")
	p(e) // strconv.Atoi: parsing "hello": invalid syntax
}
