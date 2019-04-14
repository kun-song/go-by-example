package main

import (
	"fmt"
	"os"
)

// go run xx.go 中，run 和 xx.go 即为 go 的命令行参数
func main() {
	p := fmt.Println

	// os.Args 包含所有命令行参数，第一个为 app 的全路径名
	args := os.Args

	// [1:0] 去掉首元素
	argsWithFirst := os.Args[1:]

	p(args)
	p(argsWithFirst)
	p(argsWithFirst[2])

	// [src] go build main.go                                                                                   12:24:05
	// [src] ./main a b c d                                                                                     12:34:24
	// [./main a b c d]
	// [a b c d]
	// c
}
