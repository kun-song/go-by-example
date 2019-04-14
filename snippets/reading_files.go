package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 读文件大部分函数都需要检查返回值，抽成单独的函数
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	p := fmt.Println
	path := "/tmp/data"

	// 将文件全部内容读入内存
	data, err := ioutil.ReadFile(path)
	check(err)
	p(string(data))

	// 更细粒度的控制，需要先打开文件
	f, err := os.Open(path)
	check(err)

}
