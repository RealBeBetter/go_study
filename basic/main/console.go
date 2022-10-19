package main

import "fmt"

func main() {
	var i string
	// Scanln读取一行
	fmt.Scanln(&i)
	fmt.Println("i = ", i)

	var j string
	var m float32
	var n bool

	fmt.Scanf("%d %f %s %t", &i, &m, &j, &n)
	fmt.Println("i = ", i, "j = ", j, "m = ", m, "n = ", n)

}
