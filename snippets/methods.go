package main

import (
	"fmt"
)

/**
 * Go 支持在 struct type 上定义方法，有两种 receiver type：
 *   1. pointer receiver type
 *   2. value receiver type
 */

type rect struct {
	width, height int
}

// area 方法的 receiver type 为 *rect
func (r *rect) area() int {
	return r.width * r.height
}

// receviver type 为 value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	/**
	 * pointer receiver type 与 value receiver type 使用上差别很小，pointer 有两个优点：
	 *   1. avoiding copying on method calls
	 *   2. allow the method to mutate the receiving struct
	 */
	r := rect{width: 10, height: 5}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

}
