package main

import (
	"fmt"
)

/**
 * slice 是 Go 的核心数据结构，使用远比 array 频繁，两者区别：
 *   1. array 类型由数组长度 + 元素类型确定
 *   2. slice 类型仅由元素类型确定
 */
func main() {
	/**
	 * slice 创建：make
	 */
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// set/get 与 array 类似
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 与 array 类似
	fmt.Println("len:", len(s))

	/**
	 * 除以上与 array 类似的操作外，slice 还支持一些高级操作
	 */

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// slice 支持 copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// "slice" 操作
	l := s[2:5] // 前闭后开区间
	fmt.Println("slice1:", l)

	l = s[:5]
	fmt.Println("slice2:", l)

	l = s[2:]
	fmt.Println("slice3:", l)

	// 声明 + 初始化 slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 多维 slice，与多维数组不同，内层 slice 的长度可变
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	// 2d: [[0] [1 2] [2 3 4]]
	fmt.Println("2d:", twoD)
}
