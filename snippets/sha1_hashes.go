package main

import (
	"crypto/sha1"
	"fmt"
)

/**
 * SHA1 hashes：一种摘要算法，用于对二进制 or 文本文件计算一个简短的标识符
 *
 * `crypto/*` 实现不同哈希算法，如 `crypto/sha1`、`crypto/md5`
 */
func main() {
	s := "Hello world!"

	// 1. sha1.New()
	hash := sha1.New()

	// 2. Write(target)
	hash.Write([]byte(s))

	// 3. Sum()
	bs := hash.Sum(nil) // 在结尾添加参数 []byte，一般不需要添加，所以为 nil

	fmt.Println(bs)        // [211 72 106 233 19 110 120 86 188 66 33 35 133 234 121 112 148 71 88 2]
	fmt.Printf("%x\n", bs) // d3486ae9136e7856bc42212385ea797094475802
}
