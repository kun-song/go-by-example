package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * line filter：一类程序，从 stdin 读取输入，处理，然后将结果输出到 stdout，如 sed/grep 等
 */
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// [src] go run main.go                                                                                     12:09:29
	// hello
	// HELLO
	// world
	// WORLD
}
