package main

import "fmt"

func main() {
	/**
	 * if 中的 condition 无需括号，但 {} 是必须的，即使只有一个语句
	 */
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 可以省略 else
	if 8%4 == 0 {
		fmt.Println("8 is divided by 4")
	}

	/**
	 * condition 之前可以有一个 statement，其中声明的变量在 if 分支、else 分支中可见
	 */
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has more than one digit")
	}
}
