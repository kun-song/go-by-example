package main

import (
	"fmt"
)

func main() {
	// 创建 length = 5 的 int 数组，元素默认初始化为 int 类型的 zero value
	var a [5]int
	fmt.Println("emp:", a) // 打印数组所有元素

	a[4] = 111
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len 函数
	fmt.Println("len:", len(a))

	// 声明 + 初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
