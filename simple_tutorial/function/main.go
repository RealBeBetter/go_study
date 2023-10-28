/**
 * @author wei.song
 * @since 2023/10/27 17:50
 */
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-----------function test-----------")
	fmt.Println("sum =", add(5, 10))
	divide, mod := div(100, 17)
	fmt.Println("divide =", divide, ", mod =", mod)

	fmt.Println("-------------handle exception-------------")
	openFileWithException()

	fmt.Println("-------------custom error-------------")
	if err := customError(""); err != nil {
		fmt.Println(err)
	}

	fmt.Println("-------------panic error-------------")
	// panic: runtime error: index out of range [3] with length 3
	// fmt.Println(getPanicError(3))

	fmt.Println("-------------panic error with defer-------------")
	// panic: runtime error: index out of range [3] with length 3
	fmt.Println(getPanicErrorWithDefer(3))
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

func openFileWithException() {
	// open filename.txt: no such file or directory
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println("open file error:", err)
	}
}

// error 是能够预知的错误，panic 是不可预知的错误；error 是业务逻辑的一部分，panic 不是
// 类似于 Java 中的 Exception 和 Error，Exception 是可以预知的错误，Error 是不可预知的错误
// 对于 Go 来说，则是 error 是可以预知的错误，panic 是不可预知的错误
func customError(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is empty")
	}

	fmt.Println("name is", name)
	return nil
}

func getPanicError(index int) int {
	array := [3]int{1, 2, 3}
	return array[index]
}

func getPanicErrorWithDefer(index int) (res int) {
	// panic: runtime error: index out of range [3] with length 3
	// 调用原函数，不会直接返回，而是先执行 defer 中的函数，然后再返回
	defer func() {
		if res := recover(); res != nil {
			fmt.Println("panic error:", res)
			res = -1
		}
	}()

	// 在这里使用 defer 定义了异常处理的函数，在协程退出前，会执行完 defer 挂载的任务
	// 因此如果触发了 panic，控制权就交给了 defer
	// 在 defer 的处理逻辑中，使用 recover，使程序恢复正常，并且将返回值设置为 -1
	// 在这里也可以不处理返回值，如果不处理返回值，返回值将被置为默认值 0
	return getPanicError(index)
}
