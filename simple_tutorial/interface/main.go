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
