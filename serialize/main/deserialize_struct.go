/**
 * @author: Real
 * Date: 2022/11/11 23:26
 */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 反序列化 struct
	var structSerialize = "{\"name\":\"Real\",\"age\":22}"
	// 表示用该变量接收序列化的结果
	var serializeType map[string]interface{}

	// 调用反序列化方法
	err := json.Unmarshal([]byte(structSerialize), &serializeType)
	if err != nil {
		fmt.Println("json deserialize error , error = ", err)
		return
	}

	fmt.Println(serializeType)
}
