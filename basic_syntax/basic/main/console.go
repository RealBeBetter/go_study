package main

import "fmt"

func main() {
	var i string
	// Scanln 读取一行
	_, err := fmt.Scanln(&i)
	if err != nil {
		return
	}
	fmt.Println("i = ", i)

	var j string
	var m float32
	var n bool

	scanf, err := fmt.Scanf("%d %f %s %t", &i, &m, &j, &n)
	if err != nil {
		err.Error()
		return
	}
	println(scanf)
	fmt.Println("i = ", i, "j = ", j, "m = ", m, "n = ", n)

}
