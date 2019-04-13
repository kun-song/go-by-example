package main

import (
	"fmt"
	"os"
)

/**
 * defer 推迟函数调用，通常用于清理目的，类似 Java 中的 finally
 *
 * defer 函数在其父作用域最后执行
 */
func main() {
	f := createFile("/tmp/xxx.go")

	// main 函数执行的最后，将执行 closeFile
	defer closeFile(f)

	writeFile(f, "hello world")
}

func createFile(path string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File, content string) {
	fmt.Println("writing")
	fmt.Fprintln(f, content)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
