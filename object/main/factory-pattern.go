package main

import "fmt"

func main() {
	// 工厂模式
	student := CreateStudent("Real", 22, 9)
	fmt.Println(*student)
	fmt.Println(student.GetStudentName())
}

type student struct {
	Name  string
	Age   int32
	Grade int32
}

func CreateStudent(name string, age int32, grade int32) *student {
	// 暴露函数名首字母大写的函数，重当构造方法
	return &student{name, age, grade}
}

func (stu *student) GetStudentName() string {
	return stu.Name
}
