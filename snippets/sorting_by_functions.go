package main

import (
	"fmt"
	"sort"
)

/**
 * 自定义排序规则（非自然序）：1. 自定义类型，2. 实现 sort.Interface，3. 使用 sort.Sort 排序
 */

// 1. 为通过自定义函数排序，需要先有个类型
type byLength []string

// 2. 为该类型实现 sort.Interface，即 Len/Less/Swap 三个函数，一般只有 Less 实现比较特殊，其他各类型类似
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	xs := []string{"aaaa", "bbb", "c"}

	sort.Sort(byLength(xs))
	fmt.Println(xs) // [c bbb aaaa]

	sort.Strings(xs)
	fmt.Println(xs) // [aaaa bbb c]
}
