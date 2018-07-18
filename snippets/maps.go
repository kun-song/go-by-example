package main

import (
	"fmt"
)

/**
 * map 也是 Go 的核心数据结构，类似其他语言的 hashes, dicts 等
 */
func main() {
	// 创建 map：make(map[key-type]value-type)
	m := make(map[string]int)

	// set
	m["k1"] = 10
	m["k2"] = 100
	fmt.Println(m)

	// get
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// len
	fmt.Println("len:", len(m))

	// delete
	delete(m, "k1")
	fmt.Println("map:", m)

	/**
	 * get 返回两个值，第二个代表 map 中是否存在被查询的 key，用以区分 key 不存在和 key 存在，但
	 * value 为 0 两个场景
	 */
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 声明 + 初始化
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)
}
