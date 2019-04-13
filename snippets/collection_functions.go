package main

import (
	"fmt"
	"strings"
)

/**
 * 集合操作
 *
 * Go 不支持泛型，通常要根据具体要求，自己实现集合操作
 */

func index(xs []string, t string) int {
	for i, x := range xs {
		if x == t {
			return i
		}
	}

	return -1
}

func include(xs []string, t string) bool {
	return index(xs, t) > 0
}

func any(xs []string, p func(string) bool) bool {
	// slice 指定一个参数，则返回 index，两个则为 index, value
	for _, x := range xs {
		if p(x) {
			return true
		}
	}

	return false
}

func all(xs []string, p func(string) bool) bool {
	for _, x := range xs {
		if !p(x) {
			return false
		}
	}

	return true
}

func filter(xs []string, p func(string) bool) []string {
	ys := make([]string, 0)
	for _, x := range xs {
		if p(x) {
			ys = append(ys, x) // 注意 append 返回一个新 slice
		}
	}

	return ys
}

func map2(xs []string, f func(string) string) []string {
	ys := make([]string, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

func main() {
	xs := []string{"hello", "world!", "abc"}

	fmt.Println(index(xs, "abc"))
	fmt.Println(include(xs, "world!"))

	// 匿名函数
	fmt.Println(any(xs, func(x string) bool {
		return strings.HasPrefix(x, "ab")
	}))

	fmt.Println(all(xs, func(x string) bool {
		return strings.HasPrefix(x, "ab")
	}))

	fmt.Println(filter(xs, func(x string) bool {
		return strings.Contains(x, "!")
	}))

	// 函数字面值
	fmt.Println(map2(xs, strings.ToUpper))
}
