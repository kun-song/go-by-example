package main

import (
	"fmt"
	"sort"
)

/**
 * Go 语言的 sort 包提供了对 **内置类型** 和 **自定义类型** 的排序
 */
func main() {
	// 1. 每个内置类型的排序函数不同，2. 就地排序，修改排序对象
	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{3, 2, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	// 判断是否有序
	fmt.Println(sort.IntsAreSorted(ints))
}
