package main

import "fmt"

/**
 * Go 只有一种循环结构，即 for，有 3 种基本形式的 for 循环
 */
func main() {
	/**
	 * 1. 只有一个 condition 表达式的 for 语句，等价于 while
	 */
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	/**
	 * 2. 最基本的 for
	 */
	for j := 4; j <= 7; j++ {
		fmt.Println(j)
	}

	/**
	 * 3. 没有 condition，则类似 while(true)
	 */
	for {
		fmt.Println("loop")
		break
	}

	// for 中可以使用 continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
