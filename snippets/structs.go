package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	// 初始化指定字段
	fmt.Println(person{age: 20, name: "Kyle"})

	// 初始化部分字段，其他字段被初始化为各自类型的 zero value
	fmt.Println(person{name: "Miake"})

	// & 获取指针
	fmt.Println(&person{name: "Nick", age: 21})

	// . 访问字段
	p := person{name: "Sean", age: 20}
	fmt.Println(p.age)

	// 自动 dereference
	sp := &p
	fmt.Println(sp.name)

	// struct 可变
	sp.age = 100
	fmt.Println(sp.age)

}
