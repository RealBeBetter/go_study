package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 将结构体序列化为 JSON 字符串
	personSerialize := PersonSerialize{Name: "Real", Age: 22}

	jsonStruct, err := json.Marshal(personSerialize)
	if err == nil {
		fmt.Println(jsonStruct)
		fmt.Println(string(jsonStruct))
	} else {
		fmt.Println("json serialize error, error is ", err)
		return
	}
}

// PersonSerialize
// 如果命名为大写，要想序列化字符串为小写，需要使用 json tag 名称
type PersonSerialize struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
