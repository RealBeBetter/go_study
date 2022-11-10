package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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

	// 如果文件太大，需要使用 ioUtil 读取文件
	bytes, err := ioutil.ReadFile("file/main/test.txt")
	if err != nil {
		fmt.Println("io util read file error, error = ", err)
		return
	}
	fmt.Println(string(bytes))
}
