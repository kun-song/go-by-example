package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 直接用 string 表示正则
	ifMatch, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(ifMatch) // true

	// 用编译过的 Regexp
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))           // true
	fmt.Println(r.FindString("peach first"))      // peach
	fmt.Println(r.FindStringIndex("peach first")) // [0 5]（起始 index）

	// Submatch 将匹配 p([a-z]+)ch 和 [a-z]+ 两个正则
	fmt.Println(r.FindStringSubmatch("peach first"))      // [peach ea]
	fmt.Println(r.FindStringSubmatchIndex("peach first")) // [0 5 1 3]

	// All 返回所有匹配，而非只有第一个匹配
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]（-1 所有匹配）
	fmt.Println(r.FindAllString("peach punch pinch", 1))  // [peach]（最多 1 个匹配）
	fmt.Println(r.FindAllString("peach punch pinch", 2))  // [peach punch]（最多 2 个匹配）

	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1)) // [[0 5] [6 11] [12 17]]

	// 未完待续
}
