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

type StudentField struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (stu Student) Show(name string, age int) {
	fmt.Println("Name is", name, ", age is", age)
}

func (stu Student) GetAge() int {
	return stu.Age
}

var (
	ptr    *Student
	rType  reflect.Type
	rValue reflect.Value
)

func main() {
	student := Student{"雨下一整晚", 20}
	other := StudentOther{"Real", 22}
	reflectStruct(student)
	reflectStruct(other)

	// 获取所有属性
	fmt.Println()
	reflectAttr(StudentField{"雨下一整晚Real", 22})

	// 调用结构体方法
	fmt.Println("invoke struct method...")
	reflectMethod(Student{"TempName", 18})

	// 调用结构体获取值方法
	fmt.Println("invoke get struct value method...")
	reflectStructValue(Student{"TempName", 18})

	// 调用反射构造结构体
	fmt.Println("invoke reflect to create struct...")
	reflectCreate()
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

func reflectAttr(n interface{}) {
	rType := reflect.TypeOf(n)
	rValue := reflect.ValueOf(n)

	if rValue.Kind() != reflect.Struct {
		// 如果不是Struct类别，直接结束
		return
	}

	num := rValue.NumField() // 获取字段数量
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d value = %v\n", i, rValue.Field(i)) // 获取字段值
		tagVal := rType.Field(i).Tag.Get("json")                // 获取字段的json标签值
		if tagVal != "" {
			fmt.Printf("Field %d tag = %v\n", i, tagVal)
		}
	}
}

func reflectMethod(n interface{}) {
	rValue := reflect.ValueOf(n)

	if rValue.Kind() != reflect.Struct {
		// 如果不是Struct类别，直接结束
		return
	}
	num := rValue.NumMethod() // 获取方法数量
	for i := 0; i < num; i++ {
		method := rValue.Method(i)
		fmt.Println(method) // 打印方法地址
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("Real"))
	params = append(params, reflect.ValueOf(18))
	rValue.Method(1).Call(params) // 调用方法，传入参数

	fmt.Println("...")

	res := rValue.Method(0).Call(nil) // 调用方法，接收返回值
	fmt.Println(res[0].Int())
}

func reflectStructValue(stu Student) {
	rValue := reflect.ValueOf(&stu)
	rValue.Elem().Field(0).SetString("Song Wei")
	rValue.Elem().Field(1).SetInt(22)
	fmt.Println(stu)
}

func reflectCreate() {
	ptr = &Student{}                   // 结构体指针
	rType = reflect.TypeOf(ptr).Elem() // 结构体反射类型

	rValue = reflect.New(rType) // 由结构体反射类型，获取新结构体指针反射值

	ptr = rValue.Interface().(*Student) // 把指针反射值转成空接口，并进行类型断言

	rValue = rValue.Elem() // 由结构体指针反射值获取结构体反射值

	rValue.FieldByName("Name").SetString("wei,song") // 根据属性名，对结构体反射值设置值
	rValue.FieldByName("Age").SetInt(22)

	fmt.Println(*ptr) // 输出结果
}
