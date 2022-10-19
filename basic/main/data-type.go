package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	a, sex := 2, "male"

	fmt.Printf("a的类型:%T，sex的类型:%T\n", a, sex)
	fmt.Printf("a占用内存大小：%d, sex占用内存大小：%d\n", unsafe.Sizeof(a), unsafe.Sizeof(sex))

	// 字符与字符串
	c1 := 's'
	c2 := '0'

	// 直接打印 byte，会输出 ASCII 值
	fmt.Println("c1 = ", c1, ", c2 = ", c2)
	// 输出 %c 才会变成字面量
	fmt.Printf("c1 = %c, c2 = %c\n", c1, c2)

	c3 := '宋'
	fmt.Printf("c3 = %c, 对应 unicode 码值: %d\n", c3, c3)

	// 跨行字符串
	var s = `
拜仁慕尼黑来自德甲。
它在今年欧冠八分之一淘汰赛上首回合客场3:0完胜切尔西。
`
	fmt.Println("%s\n", s)

	s1 := "abc" +
		" def" + "hij"

	fmt.Println(s1)

	a1 := 1.2
	a2 := int(a1)

	fmt.Println("a2 = ", a2)

	var i1 int32 = 12
	var i2 int8
	//var i3 int8

	i2 = int8(i1) + 127 // 运行时溢出，得不到想要结果
	//i3 = int(i1) + 128 // 直接溢出，编译错误

	fmt.Println("i2 = ", i2)

	// 基本类型转 String
	n1 := 5
	var s0 = fmt.Sprintf("%d\n", n1)
	fmt.Printf("s type:%T, s = %v\n", s0, s0)
	b := 3
	s0 = fmt.Sprintf("%t", b)
	fmt.Printf("s type:%T, s = %v\n", s0, s0)
	// 调用函数进行转换
	var num int = 12345
	var str string
	str = strconv.Itoa(num) //int 转string
	fmt.Println(str)
	s0 = strconv.FormatInt(int64(n1), 10) // 10表示十进制
	fmt.Printf("s type:%T, s = %v\n", s0, s0)
	s0 = strconv.FormatFloat(a1, 'f', 10, 64) // 'f'表示浮点数类型、10表示精度10位，64表示float64
	fmt.Printf("s type:%T, s = %v\n", s0, s0)
	s0 = strconv.FormatBool(false)
	fmt.Printf("s type:%T, s = %v\n", s0, s0)

	// string 转基本类型
	var b2, _ = strconv.ParseBool(s0) // go中可以有多个返回值，_表示接收但忽略
	fmt.Printf("b2 type:%T, b2 = %v\n", b2, b2)
	var i0, _ = strconv.ParseInt("1233", 10, 64) // 后两个参数分别表示进制和转换成int的位数
	fmt.Printf("i0 type:%T, i0 = %v\n", i0, i0)
	var f0, _ = strconv.ParseFloat("21.291", 64) //后面的参数表示转换成float的位数
	fmt.Printf("f0 type:%T, f0 = %v\n", f0, f0)

	// 指针类型
	var ptr *int = &num // ptr是指针变量  ptr的类型为 *int   值为&num
	fmt.Println("ptr 的指针值：%x", ptr)
	i := 1
	ptr0 := &i
	fmt.Printf("%x, %d, %x", ptr0, *ptr0, &ptr0) //*ptr0表示取出ptr0指向地址对应的值

	(*ptr0) = (*ptr0) * 10
	fmt.Printf("%v\n", i)
}
