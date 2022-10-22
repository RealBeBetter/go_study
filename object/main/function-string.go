package main

import "fmt"

func main() {
	p := new(PersonString)
	(*p).name = "Real"
	(*p).age = 20
	(*p).address = "Hunan"

	fmt.Println(*p)
	fmt.Println(p)
	toString := p.toString()
	fmt.Println(toString)

	p1 := PersonString{age: 21, name: "Virtual", address: "Jiangsu"}
	fmt.Println(p1)
	str := p1.string()
	fmt.Println(str)

	// 调用方式不一样 函数： 函数名(实参列表) 方法： 变量名.方法名(实参列表)
	// 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	// 对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反之亦然
}

// 简单实现 java 中的 toString 方法
func (p PersonString) toString() string {
	return fmt.Sprintf("name: %s, age: %d, address: %s", p.name, p.age, p.address)
}

func (p *PersonString) string() string {
	return fmt.Sprintf("name: %s, age: %d, address: %s", p.name, p.age, p.address)
}

type PersonString struct {
	age     int
	name    string
	address string
}
