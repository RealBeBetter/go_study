package main

import "fmt"

//func (t type) methodName (参数列表) (返回值列表){
//	方法体
//	return 返回值
//}
//// t type 表示这个方法和type这个类型进行绑定 t为type的一个实例

func main() {
	person0 := PersonStruct{"Bob", 23, "California"}

	person1 := new(PersonStruct)
	(*person1).Name = "Jason"
	(*person1).Age = 24
	// 方法调用
	person0.test()
	(*person1).test()

	// 方法中的修改是否会同步到外面
	m := make(map[string]int, 5)
	m["Chinese"] = 79
	person2 := PersonStruct1{"Real", 20, "Hunan Changsha", m}

	person3 := new(PersonStruct1)
	(*person3).Name = "Jason"
	(*person3).Age = 21
	m1 := make(map[string]int, 5)
	m1["Math"] = 90
	(*person3).Scope = m1

	fmt.Println("init...person2: ", person2)
	fmt.Println("init...person3: ", *person3)
	// 初始化完毕，测试方法之间会不会产生影响
	person2.test1()
	fmt.Println("test1...person2: ", person2)
	person3.test1()
	fmt.Println("test1...person3: ", *person3)
	// 查看 Map 的值
	fmt.Println("m: ", m, "m1: ", m1)

	// 为了修改更加高效，会使用指针传参的方法调用
	person4 := person3
	test2 := person4.test2(10)
	fmt.Println("test2: ", test2)
	fmt.Println("test2...person4: ", *person4)
	person5 := person3
	testAdd := (*person5).test2(10)
	fmt.Println("testAdd: ", testAdd)
	fmt.Println("testAdd...person5: ", *person5)

	// 调用别名方法
	var num integer
	num = 8
	num.test(6)
	fmt.Println("num: ", num)
}

type PersonStruct struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Hometown string `json:"homeTown"`
}

// 表示 test 方法是和 PersonStruct 类型进行的一个绑定
// 即：只有 PersonStruct 对象才能调用改方法，相当于属于特定类的一个方法
func (p PersonStruct) test() {
	fmt.Println("name:", p.Name, "\tage:", p.Age, "\thometown:", p.Hometown)
}

type PersonStruct1 struct {
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Hometown string         `json:"homeTown"`
	Scope    map[string]int `json:"scope"`
}

func (p PersonStruct1) test1() {
	p.Age += 1
	p.Scope["Chinese"] += 1
}

func (p *PersonStruct1) test2(n int) string {
	(*p).Scope["Math"] += n
	(*p).Scope["Chinese"] += n
	return "success"
}

// 实现方法，要先定义别名
type integer int

func (i *integer) test(n int) {
	// int和integer虽然只是别名关系，但依旧不是同一个类型
	*i += integer(n)
}
