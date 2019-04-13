package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
 * Go 内置 JSON 解析支持（内置类型 + 自定义类型）
 */

type response1 struct {
	Page   int
	Furits []string
}

type response2 struct {
	Page   int      `json:"page"` // 自定义序列化后的 key 名字
	Furits []string `json:"fruits"`
}

var p = fmt.Println

func main() {
	////////////////// 序列化 ///////////////////

	// bool -> JSON byte[]/string
	bolB, _ := json.Marshal(true)
	p(bolB)         // [116 114 117 101]
	p(string(bolB)) // true

	// int -> JSON byte[]/string
	intB, _ := json.Marshal(10)
	p(intB)         // [49 48]
	p(string(intB)) // 10

	// float -> JSON
	fltB, _ := json.Marshal(10.5)
	p(fltB)         // [49 48 46 53]
	p(string(fltB)) // 10.5

	// string -> JSON
	strB, _ := json.Marshal("Hello")
	p(strB)         // [34 72 101 108 108 111 34]
	p(string(strB)) // "Hello"

	// slice -> JSON
	sliB, _ := json.Marshal([]string{"hello", "world"})
	p(sliB)         // [91 34 104 101 108 108 111 34 44 34 119 111 114 108 100 34 93]
	p(string(sliB)) // ["hello","world"]

	// map -> JSON
	mapB, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	p(mapB)         // [123 34 97 34 58 49 44 34 98 34 58 50 125]
	p(string(mapB)) // {"a":1,"b":2}

	// 自定义类型 -> JSON，默认以字段名为 key
	res1, _ := json.Marshal(&response1{
		Page:   10,
		Furits: []string{"apple", "peach"}})
	p(res1)         // [123 34 80 97 103 101 34 58 49 48 44 34 70 117 114 105 116 115 34 58 91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 93 125]
	p(string(res1)) // {"Page":10,"Furits":["apple","peach"]}

	// 自定义 key 名字
	res2, _ := json.Marshal(&response2{
		Page:   10,
		Furits: []string{"apple", "peach"}})
	p(string(res2)) // {"page":10,"fruits":["apple","peach"]}

	////////////////// 反序列化 ///////////////////

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// [123 34 110 117 109 34 58 54 46 49 51 44 34 115 116 114 115 34 58 91 34 97 34 44 34 98 34 93 125]
	p(byt)

	// dat will hold a map of strings to arbitrary data types
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) // map[strs:[a b] num:6.13]

	// 反序列化自定义类型
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	resp2 := response2{}
	if err := json.Unmarshal([]byte(str), &resp2); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp2)   // {Page:1 Furits:[apple peach]}
	fmt.Println(resp2.Furits[0]) // apple

	// 重定向序列化结果
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}

	// 注意，没有显式调用 PrintX，但我们把 Encoder 的输出设置为 os.Stdout，所以在终端会输出
	enc.Encode(d) // {"apple":5,"lettuce":7}
}
