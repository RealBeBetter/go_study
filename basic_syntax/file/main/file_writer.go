package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

const (
	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

func main() {
	// 文件写入操作
	filePath := "file/main/writer_test.txt"

	// 判断文件或文件夹是否存在使用os.Stat()返回错误值进行判断：
	// 如果返回错误为nil说明文件或文件夹存在
	// 如果返回的错误类型使用os.IsNotExist()判断为true，说明不存在
	// 如果返回的错误为其它类，则不确定是否存在
	stat, err := os.Stat(filePath)
	if err == nil {
		fmt.Println("The file/path already exists.")
	}
	fmt.Println(stat)

	// 如果打开已存在的文件，需要清空内容，就要把模式换成os.O_TRUNC
	// 如果打开已存在的文件，需要写入文件，可以使用模式os.O_RDWR
	file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writeString, err := writer.WriteString("New content " + fmt.Sprintf("%d", i) + " ....\n")
		if err != nil {
			fmt.Println("Write file error, error = ", err)
			return
		}
		fmt.Println("Write file content: ", writeString)
	}

	err = writer.Flush()
}
