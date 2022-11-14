/**
 * @author: Real
 * Date: 2022/11/12 13:20
 */
package main

import "fmt"

func main() {
	// WithGroup 可以等待对应的协程执行完再往下执行，可以让主程序等待协程执行之后再执行

}

func show(i int) {
	fmt.Println("show method running...", i)
}
