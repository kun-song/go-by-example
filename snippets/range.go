package main

import (
	"fmt"
)

/**
 * range：遍历多种数据结构
 */
func main() {
	// 1. 遍历 slice（array)
	xs := []int{1, 2, 3, 4}
	sum := 0
	for _, x := range xs {
		sum += x
	}
	fmt.Println("sum =", sum)

	// 遍历 slice/array 时，会同时获得 index 和 value
	for i, x := range xs {
		if x == 3 {
			fmt.Println("index =", i)
		}
	}

	// 2. 遍历 map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 只遍历 map 的 key
	for k := range kvs {
		fmt.Println("key =", k)
	}

	// 3. 遍历 string
	for i, c := range "Hello" {
		fmt.Println(i, c)
	}
}
