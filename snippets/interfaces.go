package main

import (
	"fmt"
	"math"
)

/**
 * interfaces are named collections of method signatures.
 */

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 要实现一个 interface，只需要实现该 interface 定义的所有方法即可

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 若变量类型为 interface，则可以用该变量调用该 interface 上的方法
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// rect 和 circle 都实现了 geometry 定义的方法，因此可以视为 geometry
func main() {
	r := rect{width: 20, height: 5}
	c := circle{radius: 5}

	// {20 5}
	// 100
	// 50
	measure(r)
	// {5}
	// 78.53981633974483
	// 31.41592653589793z
	measure(c)
}
