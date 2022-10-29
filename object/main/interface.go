package main

import "fmt"

func main() {

	//type 接口名 interface{
	//	method1(参数列表) 返回值列表
	//	method2(参数列表) 返回值列表
	//}

	c0 := C{}
	d0 := D{}
	e0 := E{}

	i := c0.add(8, 9)
	fmt.Println(i)

	j := d0.add(5, 6)
	fmt.Println(j)

	e0.sub(c0)

}

type ICalculate interface {
	add(a int, b int) int
	sub(a int, b int) int
	addNone()
	subNone()
}

// C 定义两个结构体，实现该接口
type C struct {
}

func (c C) sub(a int, b int) int {
	fmt.Println("sub method")
	return a + b
}

func (c C) addNone() {
	fmt.Println("addNone...")
}

func (c C) subNone() {
	fmt.Println("subNone...")
}

type D struct {
}

func (d D) add(a int, b int) int {
	fmt.Println("D..add")
	return a + b
}

func (c C) add(a int, b int) int {
	fmt.Println("C...add")
	return a + b
}

type E struct {
}

// 接口是引用类型，所以这里传递的是变量的引用
func (e *E) add(ic ICalculate) {
	ic.addNone()
}

func (e *E) sub(ic ICalculate) {
	ic.subNone()
}
