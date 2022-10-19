package main

import "fmt"

func main() {

	// Go中没有while，do…while循环，Go语言支持goto跳转，但不支持使用
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break // break  默认会跳出最近的for循环
			}
			fmt.Println("j=", j)
		}
	}

	var a int = 5
	var b int = 6

	swap(&a, &b)
	fmt.Println(a, b)

}

func swap(ptr1 *int, ptr2 *int) {
	var temp *int = ptr1
	ptr1 = ptr2
	ptr2 = temp
}
