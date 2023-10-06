package main

import (
	"fmt"
	"os"
)

func main() {
	// go 中的文件打开与关闭
	file, err := os.Open("file/main/test.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
		return
	}

	fmt.Println("file open success , file = ", file)

	err = file.Close()
	if err != nil {
		fmt.Println("file close error = ", err)
	}
}
