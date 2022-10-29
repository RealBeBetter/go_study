package main

import "fmt"

func main() {
	i := integer0(1)
	fmt.Println(i)

	var cals []ICalculate
	// 接口数组添加新的接口
	cals = append(cals)
	fmt.Println(cals)

}

type integer0 int

func (i integer0) add() {
	fmt.Println("integer..add")
}
