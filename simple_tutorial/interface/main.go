/**
 * @author Real
 * @since 2023/10/28 15:16
 */
package main

import "fmt"

func main() {
	fmt.Println("-----------interface test-----------")
	student := new(Student)
	student.name = "Real"
	student.age = 18
	// hello Real
	println(student.getName())

	fmt.Println("-----------interface test-----------")
	teacher := &Teacher{age: 25, name: "Tom", course: "Math"}
	println(teacher.getName())

	fmt.Println("-----------interface test-----------")
	// define interface variable
	var person Person = student
	println(person.getName())

	fmt.Println("-----------interface test-----------")
	// define interface variable
	var person2 Person = teacher
	println(person2.getName())

	// struct need implement all interface methods, or else will throw error

	fmt.Println("-----------interface test-----------")
	var _ Person = (*Student)(nil)
	fmt.Println("you can use this way to check whether struct implement interface")
	var _ Person = (*Teacher)(nil)
	// you can use this way to check whether struct implement interface

	fmt.Println("-----------interface test-----------")
	// interface type convert to struct type
	stu := person.(*Student)
	println(stu.name)
	// this will cause
	// panic: interface conversion: main.Person is *main.Student, not *main.Teacher
	// teacherConvert := person.(*Teacher)
	// println(teacherConvert.name)

	fmt.Println("-----------interface test-----------")
	emptyInterface()
}

type Person interface {
	// define interface method
	getName() string
}

type Student struct {
	// define struct field
	name string
	age  int
}

// implement interface method
func (stu *Student) getName() string {
	return stu.name
}

type Teacher struct {
	// define struct field
	name   string
	age    int
	course string
}

func (teacher *Teacher) getName() string {
	return teacher.name
}

func emptyInterface() {
	// key = string, value = any type
	m := make(map[string]interface{})

	// 一定程度上，interface{} 类型可以代替 Java 中的 Object 类型
	m["name"] = "Real"
	m["age"] = 18
	m["course"] = []string{"Math", "English"}
	m["score"] = map[string]int{"Math": 100, "English": 90}

	fmt.Println(m)
}
