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

	/**
	https://blog.csdn.net/qq_39697306/article/details/115085149
	如果需要使用在同一个包中定义的其他文件中的结构体，需要将项目放到 $GOPATH/src 下
	*/
	// 同一个目录下面不能有个多 package main ，此处需要重新定义
	var serializePerson PersonSerialize

	// 调用反序列化方法
	err := json.Unmarshal([]byte(structSerialize), &serializePerson)
	if err != nil {
		fmt.Println("struct deserialize error , error = ", err)
		return
	}

	fmt.Println(serializePerson)
}

// PersonSerialization
// 如果命名为大写，要想序列化字符串为小写，需要使用 json tag 名称
//type PersonSerialization struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
