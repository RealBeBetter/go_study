package main

import "fmt"

func main() {

	graduate := &Graduate{}
	graduate.Age = 20
	graduate.Major = "软件工程"
	graduate.Name = "Real"
	fmt.Println(*graduate)

}

type Student struct {
	Name string
	Age  int
}

type Graduate struct {
	// 匿名结构体
	Student
	Major string
}

func (student *Student) getName() string {
	return student.Name
}

func (graduate *Graduate) getMajor() string {
	return graduate.Major
}
