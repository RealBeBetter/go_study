/**
 * @author Real
 * @since 2023/10/28 16:03
 */
package main

import "fmt"

func main() {
	fmt.Println("-----------module test-----------")
	// 同一个 package 下的 function 可以直接引用
	fmt.Println(calcAdd(1, 2))
}
