package main

import "fmt"

/**
 * 变量声明两种方式：
 *   1. var （1）声明 + 可选的初始化器  （2）可多个变量
 *   2. :=  （1）声明 + 初始化器       （2）可多个变量
 *
 * := 仅仅是 *var 声明 + 初始化* 的语法糖
 */
func main() {
	// 1. 可从初始化器自动推断变量的类型
	var a = "initial"
	fmt.Println(a)

	// 2. 同时声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	/**
	 * 若无初始化器，则
	 *   1. 必须显式声明类型
	 *   2. 默认初始化为该类型的 zero value
	 */
	var e int
	fmt.Println(e)

	// :=
	f := "short hand"
	fmt.Println(f)
}
