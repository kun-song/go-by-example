package main

import (
	"fmt"
	"strings"
)

/**
 * `strings` 包：字符串相关函数
 */

// 使用 p 替代 fmt.Println
var p = fmt.Println

func main() {
	p("Contains: ", strings.Contains("Hello", "H")) // true
	p("Count: ", strings.Count("Hello", "l"))       // 2

	p("HasPrefix: ", strings.HasPrefix("Hello", "He")) // true
	p("HasSuffix: ", strings.HasSuffix("Hello", "lo")) // true

	p("Index: ", strings.Index("Hello", "l")) // 2

	p("Join: ", strings.Join([]string{"a", "b"}, "-")) // a-b
	p("Repeat: ", strings.Repeat("x", 5))              // xxxxx

	p("Replace: ", strings.Replace("Hello", "l", "L", 1))  // HeLlo
	p("Replace: ", strings.Replace("Hello", "l", "L", -1)) // HeLLo

	p("Split: ", strings.Split("a-b-c", "-")) // [a b c]

	p("ToUpper: ", strings.ToUpper("Hello")) // HELLO
	p("ToLower: ", strings.ToLower("Hello")) // hello

	p("len: ", len("Hello")) // 5
	p("Char: ", "Hello"[1])  // 101
}
