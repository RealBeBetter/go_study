/**
 * @author: Real
 * Date: 2022/11/15 0:04
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 基本数据类型进行反射

	reflectBase(10)
}

func reflectBase(n int) {
	typeOf := reflect.TypeOf(n)
	fmt.Println("Type is", typeOf)
	fmt.Println("The name of Type is", typeOf.Name())

	valueOf := reflect.ValueOf(n)
	fmt.Printf("Value is %d, the type of value is %T\n", valueOf, valueOf)

	// 获取反射持有的整型值
	i := valueOf.Int() + 3
	fmt.Println("The value of i is", i)

	inter := valueOf.Interface()
	// 类型断言
	num := inter.(int)
	fmt.Println("num =", num)
}
