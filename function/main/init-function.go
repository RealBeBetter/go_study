package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*func 函数名 (参数列表) （返回值列表） { //返回值只有一个时可以不写（）
		//函数体
		return 返回值列表
	}*/

	fmt.Println(generateRandom(time.Now().Unix(), 100))
}

func generateRandom(time int64, _range int) int {
	rand.Seed(time)
	return rand.Intn(_range)
}

/**
源文件执行流程：全局变量定义 -> init -> main
如果此文件还引入了别的文件，就先执行被引用文件的变量定义和init
*/
func init() {
	fmt.Println("init variable_advanced..")
}
