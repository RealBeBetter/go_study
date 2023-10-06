package main

import "fmt"

const (
	KB float64 = 1 << (10 * iota) //iota是 const 结构里面，定义常量行数的索引器，每个 const 里面，iota 都从 0 开始
	MB                            //下面是一个省略调用，继承了上面的表达式
	GB
	TB
	PB
)

func main() {
	var a int = 60 //二进制是：111100
	var b int = 13 //二进制是：001101

	fmt.Printf("%b %d\n", a&b, a&b) //二进制是：1100,对应的十进制是12。说明&进行的是上下对应位的与操作
	fmt.Printf("%b %d\n", a|b, a|b) //二进制是：111101,对应的十进制是61。说明&进行的是上下对应位的或操作
	fmt.Printf("%b %d\n", a^b, a^b) //二进制是：110001,对应的十进制是49。^位运算符是上下对应位不同时，值为1

	fmt.Printf("1MB = %vKB\n", MB)
	fmt.Printf("1GB = %vKB\n", GB)
	fmt.Printf("1TB = %vKB\n", TB)
	fmt.Printf("1PB = %vKB\n", PB)

	var i int = 4
	var j int32
	var k float32
	var ptr *int

	fmt.Printf("a 变量类型为 = %T\n", i) //输出变量类型%T
	fmt.Printf("b 变量类型为 = %T\n", j)
	fmt.Printf("c 变量类型为 = %T\n", k)

	ptr = &a
	fmt.Printf("a 的内存地址为 = %p", ptr) //go里面的内存块地址通常都是用十六进制表示的，因此输出：0x10414020a
	fmt.Printf("*ptr 为 %d\n", *ptr)  //这是个指向a的内存地址的指针，因此输出：4

	// 优先级     运算符
	// 7      ^ !
	// 6      * / % << >> & &^
	// 5      + - | ^
	// 4      == != < <= >= >
	// 3      <-
	// 2      &&
	// 1      ||
}
