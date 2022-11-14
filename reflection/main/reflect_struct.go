/**
 * @author: Real
 * Date: 2022/11/15 0:18
 */
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type StudentOther struct {
	Name string
	Age  int
}

func main() {
	student := Student{"雨下一整晚", 20}
	other := StudentOther{"Real", 22}
	reflectStruct(student)
	reflectStruct(other)
}

func reflectStruct(n interface{}) {
	rType := reflect.TypeOf(n)
	rValue := reflect.ValueOf(n)

	iv := rValue.Interface()

	fmt.Println("rType = ", rType, ", iv = ", iv)
	// rType =  main.Student_rf , iv =  {szc 23}
	fmt.Printf("Type of iv = %T\n", iv)
	// Type of iv = main.Student_rf

	// 类型断言
	switch iv.(type) {
	case Student:
		student := iv.(Student)
		fmt.Println(student.Name, ", ", student.Age)
	case StudentOther:
		student := iv.(StudentOther)
		fmt.Println(student.Name, ", ", student.Age)
	}

	fmt.Println("rKind = ", rValue.Kind(), ", rKind = ", rType.Kind())
	// rKind =  struct , rKind =  struct
}
