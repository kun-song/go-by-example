package main

import (
	b64 "encoding/base64"
	"fmt"
)

/**
 * Go 支持两种 base64，两者有细微区别：
 *   + base64.StdEncoding
 *   + base64.URLEncoding
 *
 * 例如下例中，StdEncoding 和 URLEncoding 的编码结果分别为：
 *   + YWJjMTIzIT8kKiYoKSctPUB+
 *   + YWJjMTIzIT8kKiYoKSctPUB-
 *
 * 两者只有结尾不同，但都可以解码为原字符串
 */
func main() {
	p := fmt.Println

	data := "abc123!?$*&()'-=@~"

	//////////////// StdEncoding ////////////////

	// 编码
	a := b64.StdEncoding.EncodeToString([]byte(data))
	p(a) // YWJjMTIzIT8kKiYoKSctPUB+

	// 解码
	aa, _ := b64.StdEncoding.DecodeString(a)
	p(aa)         // [97 98 99 49 50 51 33 63 36 42 38 40 41 39 45 61 64 126]
	p(string(aa)) // abc123!?$*&()'-=@~

	//////////////// StdEncoding ////////////////

	// 编码
	b := b64.URLEncoding.EncodeToString([]byte(data))
	p(b) // YWJjMTIzIT8kKiYoKSctPUB-

	// 解码
	bb, _ := b64.URLEncoding.DecodeString(b)
	p(bb)         // [97 98 99 49 50 51 33 63 36 42 38 40 41 39 45 61 64 126]
	p(string(bb)) // abc123!?$*&()'-=@~
}
