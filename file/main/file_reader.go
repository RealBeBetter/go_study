package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file/main/test.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
		return
	}

	// 读取文件
	reader := bufio.NewReader(file)

	for {
		strLine, err := reader.ReadString('\n')
		if err == nil {
			// 表示文件读取成功，直接输出
			fmt.Println(strLine)
		} else if err == io.EOF {
			// 表示读取到文件末尾，结束读取
			fmt.Println("read file end")
			break
		} else {
			// 表示读取文件出现错误，打印错误
			fmt.Println("read line error, error = ", err)
		}
	}
}
