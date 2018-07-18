package main

import (
	"fmt"
	"time"
)

func main() {
	// 基本
	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 用 , 分割多个条件
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/**
	 * switch 后面没有表达式时，可以认为是一种 if/else 形式
	 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before moon")
	default:
		fmt.Println("It's after moon")
	}

	/**
	 * swith 可以比较 value，也可以比较 type
	 */
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(111)
	whatAmI("Hello")
}
