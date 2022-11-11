/**
 * @author: Real
 * Date: 2022/11/11 23:58
 */
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 反序列slice
	var sliceSerialize string = "[{\"Age\":22,\"Name\":\"Real1\"},{\"Age\":22,\"Name\":\"Real2\"},{\"Age\":22,\"Name\":\"Real3\"},{\"Age\":22,\"Name\":\"Real4\"},{\"Age\":22,\"Name\":\"Real5\"}]"

	slice := make([]map[string]interface{}, 0)

	err := json.Unmarshal([]byte(sliceSerialize), &slice)
	if err != nil {
		fmt.Println("deserialize slice failed, error is ", err)
		return
	}

	fmt.Println("deserialize slice succeed , result is : ", slice)
}
