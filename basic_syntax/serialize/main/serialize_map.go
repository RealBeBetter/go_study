/**
 * @author: Real
 * Date: 2022/11/11 23:07
 */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 对映射进行序列化操作
	// 键是空接口，表示能接收到任意类型
	var personMap map[string]interface{}
	personMap = make(map[string]interface{})

	personMap["name"] = "RealBeBetter"
	personMap["age"] = 22

	mapSerialize, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("map serialize error, error = ", err)
		return
	} else {
		fmt.Println("serialize result is ", string(mapSerialize))
		// 序列化之后的数据返回值是一个 byte 数组
		fmt.Println("serialize byte is ", mapSerialize)
	}
}
