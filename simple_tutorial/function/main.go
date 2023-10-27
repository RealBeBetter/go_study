/**
 * @author wei.song
 * @since 2023/10/27 17:50
 */
package main

import "fmt"

func main() {
	fmt.Println("-----------function test-----------")
	fmt.Println("sum =", add(5, 10))
	divide, mod := div(100, 17)
	fmt.Println("divide =", divide, ", mod =", mod)
}

// the definition of function

//func funcName(param1 Type1, param2 Type2, ...) (return1 Type3, ...) {
//	// body
//}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}
