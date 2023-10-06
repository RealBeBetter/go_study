/**
 * @author: Real
 * Date: 2022/11/15 0:42
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999") // 和本地9999端口建立tcp连接
	if err != nil {
		fmt.Println("Connect to server failure..")
		return
	}

	fmt.Println("Connected to server whose ip is ", conn.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin) // 建立控制台的reader

	for {
		line, err := reader.ReadString('\n') // 读取控制台一行信息
		if err != nil {
			fmt.Println("Read String error :", err)
		}

		line = strings.Trim(line, "\r\n")

		if line == "quit" {
			break
		}

		_, err = conn.Write([]byte(line)) // 向服务端发送信息，返回发送的字节数和错误
		if err != nil {
			fmt.Println("Write to server error:", err)
		}
	}
}
