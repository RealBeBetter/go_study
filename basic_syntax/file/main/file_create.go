package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件创建操作
	filePath := "file/main/writer_test.txt"
	// 参数1：文件路径，参数2：打开模式，参数3：权限控制（windows无效）
	// os.OpenFile(name string, flag int, perm FileMode)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	err = file.Close()

	if err != nil {
		fmt.Println("Close file error: ", err)
		return
	}
}
