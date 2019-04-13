package main

import "fmt"

/**
 * fmt.Println 提供类似 C 语言 printf 的字符串格式化功能
 */

type point struct {
	x, y int
}

func main() {
	//////// struct 打印 /////////
	p := point{5, 10}

	fmt.Printf("%v\n", p)  // {5 10}
	fmt.Printf("%+v\n", p) // {x:5 y:10}，带字段名
	fmt.Printf("%#v\n", p) // main.point{x:5, y:10}
	// T 打印值类型
	fmt.Printf("%T\n", p) // main.point

	fmt.Printf("%t\n", true) // true

	//////// 整数打印 /////////
	fmt.Printf("%d\n", 123) // 十进制打印：123
	fmt.Printf("%b\n", 123) // 二进制打印：1111011
	fmt.Printf("%x\n", 123) // 十六进制打印：7b
	fmt.Printf("%c\n", 123) // 打印数字对应的字符：{

	//////// 浮点数打印 /////////
	fmt.Printf("%f\n", 12.3) // 十进制：12.300000
	fmt.Printf("%e\n", 12.3) // 科学计数法 1：1.230000e+01
	fmt.Printf("%E\n", 12.3) // 科学计数法 2：1.230000E+01

	// 未完待续
}
