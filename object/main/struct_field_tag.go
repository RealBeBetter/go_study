package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	person := PersonJson{"Jack", 20, "北京海淀区1号"}
	personJson, err := json.Marshal(person)
	if err == nil {
		fmt.Println(personJson)
		fmt.Println(string(personJson))
	}

}

type PersonJson struct {
	// 类型后面添加的就是标签，方便实现json序列化
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
