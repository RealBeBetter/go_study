package main

import "fmt"

func main() {
	const tax int = 1
	const name_ = "szc"
	const number = 4 / 2

	//    const b_ = getVal() 	// 编译期值不确定
	//    num := 1
	//    const b_ = num / 2  	// 编译期值不确定
	//    const tax_0 int 	// 必须赋初值
	//    tax = 2	 // 不可修改

	const (
		// b、c可以不写 iota ，但是 a 必须写
		a    = iota       //0
		b    = iota       //1
		c    = iota       //2
		d, e = iota, iota //3 3
	)
	fmt.Println(a, b, c, d, e) // 0 1 2 3 3 依次加1

}
