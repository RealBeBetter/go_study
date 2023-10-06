package main

import "fmt"

type Person struct {
	age    int
	name   string
	gender int
}

func main() {
	// 1、结构体使用
	person1 := Person{20, "Real", 1}
	fmt.Println("person1: ", person1)

	person2 := person1
	fmt.Println("person2: ", person2)

	// 重新赋值使用，判断拷贝类型
	person1.name = "Virtual"
	// 修改原有的 person1 ，并不影响 person2，属于是深拷贝
	fmt.Println("after modified, person1: ", person1)
	fmt.Println("after modified, person2: ", person2)

	// 2、指针方式一声明结构体
	person3 := new(Person)
	(*person3).age = 20
	(*person3).name = "Virtual"
	(*person3).gender = 0
	fmt.Println("结构体 person3: ", *person3)

	// 3、指针方式二声明结构体
	person4 := &Person{20, "Choker", 1}
	fmt.Println("结构体 person4: ", *person4)

	rectangle := Rectangle{Point{1, 1}, Point{2, 2}}
	fmt.Println("长方形 rectangle: ", rectangle)
	fmt.Printf("%p ", &rectangle)
	fmt.Println(&rectangle.leftUp.x, &rectangle.leftUp.y, &rectangle.rightDown.x, &rectangle.rightDown.y)

	rect1 := &Rect_{leftUp: &Point{1, 1}, rightDown: &Point{3, 3}}
	fmt.Printf("%p %p\n", rect1.leftUp, rect1.rightDown)
	fmt.Println(&rect1.leftUp.x, &rect1.leftUp.y, &rect1.rightDown.x, &rect1.rightDown.y)
}

type Rect_ struct {
	leftUp, rightDown *Point
}

type Point struct {
	x, y int
}

type Rectangle struct {
	leftUp, rightDown Point
}
