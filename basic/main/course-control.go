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

	/*if 表达式1 { // 表达式可以写()，官方不推荐写
		 //代码块
	} else if 表达式2{ // else if 可省略，不能换行写
		 //代码块
	} else { // else 可省略，不能换行写
		 //代码块
	}*/

	i := 5
	if i == 1 {
		fmt.Println("第一层", i)
	} else if i == 2 {
		fmt.Println("第二层", i)
	} else {
		fmt.Println("第三层", i)
	}

	/*switch 表达式 {
		 case 表达式1，表达式2，…… :
		 // 语句块1
		 case 表达式3，表达式4，…… :
		 // 语句块2
		 // 多个case，结构同上
		 default :
		 // 语句块3
	}*/

	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("case default")
	}

	/*for 循环变量初始化 ;循环条件 ;循环变量迭代 {
		 //循环操作
	}*/

	// Go中没有while，do…while循环，Go语言支持goto跳转，但不支持使用
	for i := 0; i < 10; i++ {
		fmt.Println("for循环", i)
	}

	str := "雨下一整晚Real"
	for index, s := range str {
		fmt.Printf("%d --- %c\n", index, s)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break // break  默认会跳出最近的for循环
			}
			fmt.Println("j=", j)
		}
	}

label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1 // 一旦触发直接终止i的循环
			}
			fmt.Println("跳出标签 j=", j)
		}
		fmt.Println("跳出标签 i=", i)
	}

}
