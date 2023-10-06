/**
 * @author: Real
 * Date: 2022/11/11 23:15
 */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 切片序列化
	var slice []map[string]interface{}
	slice = make([]map[string]interface{}, 0)

	for i := 1; i <= 5; i++ {
		var temp map[string]interface{}
		temp = make(map[string]interface{}, 2)

		temp["Name"] = "Real" + fmt.Sprintf("%d", i)
		temp["Age"] = 22

		slice = append(slice, temp)
	}

	// 构建切片完毕，开始进行序列化
	serializeSlice, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("serialize slice error , error = ", err)
		return
	}

	fmt.Println("serialize slice succeed, result: ", string(serializeSlice))
	fmt.Println(serializeSlice)

}
