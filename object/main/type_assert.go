package main

import "fmt"

func main() {
	// 由于接口是一般类型，不知道具体类型，如果要转成具体类型就需要使用类型断言
	// 如果类型不匹配会报错
	var cals []ICal
	// 接口数组添加新的接口
	cals = append(cals)

	cals = append(cals, Assert{})

	// 使用方法：待断言变量.(断言类型)
	b1, succeed := cals[0].(Assert)
	if succeed {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println(b1)

	c1, success := cals[1].(Assert)
	if success {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println(c1)

	switch cals[0].(type) {
	case Assert:
		fmt.Println("type assert")
	default:
		fmt.Println("type unknown")
	}
}

type ICal interface {
	addCal(a int, b int) int
	subCal(a int, b int) int
	addCalNone()
	subCalNone()
}

type Assert struct {
}

func (assert Assert) addCal(a int, b int) int {
	fmt.Println("addCal method")
	return a + b
}

func (assert Assert) subCal(a int, b int) int {
	fmt.Println("subCal method")
	return a - b
}

func (assert Assert) addCalNone() {
	fmt.Println("addCalNone...")
}

func (assert Assert) subCalNone() {
	fmt.Println("subCalNone...")
}
