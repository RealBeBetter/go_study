/**
 * @author Real
 * @since 2023/10/28 15:08
 */
package main

import fm "fmt"

func main() {
	fm.Println("-----------struct test-----------")
	// 1. define struct
	stu := &Student{name: "Real", age: 18}
	// 2. call function
	// hello Tom, my name is Real, I am 18 years old
	fm.Println(stu.hello("Tom"))

	fm.Println("-----------struct test-----------")
	// 1. define struct
	stu2 := new(Student)
	// without set value, the default value is empty string and 0
	//stu2.name = "Real"
	//stu2.age = 18
	// 2. call function
	// hello Alice, my name is Real, I am 18 years old
	fm.Println(stu2.hello("Alice"))
}

type Student struct {
	name string
	age  int
}

// the function is defined in the struct
// 代表函数的接收者是 Student 类型的变量，需要通过 Student 类型的变量来调用 hello() 方法
func (stu *Student) hello(person string) string {
	return fm.Sprintf("hello %s, my name is %s, I am %d years old", person, stu.name, stu.age)
}
