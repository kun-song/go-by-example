package main

import (
	"errors"
	"fmt"
)

/**
 * Go 习惯把 error 通过返回值的形式返回
 */

// error 一般作为最后一个返回值，且类型为 error（内置 interface）
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// myError 自定义 error，实现 Error() 方法即可
type myError struct {
	arg  int
	prob string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &myError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func main() {

	for _, x := range []int{12, 42} {
		if v, e := f1(x); e != nil { // 惯用法
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", v)
		}
	}

	for _, x := range []int{12, 42} {
		if v, e := f2(x); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", v)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*myError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
