/**
 * @author: Real
 * Date: 2022/11/15 0:37
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	// 以tcp为例，服务端建立监听套接字，然后阻塞等待客户端连接。
	// 客户端连接后，开启协程处理客户端。
	fmt.Println("Server on..")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	// 建立tcp的监听套接字，监听本地9999号端口
	if err != nil {
		fmt.Println("Server listen error..")
		return
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
	}(listen)

	for {
		fmt.Println("Waiting for client to connect..")
		conn, err := listen.Accept() // 等待客户端连接

		if err != nil {
			fmt.Println("Client connect error..")
			continue
		}

		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				panic(err)
			}
		}(conn)

		fmt.Println("Connection established with ip:", conn.RemoteAddr().String()) // 获取远程地址
		go processClient(conn)

	}
}

func processClient(connection net.Conn) {
	for {
		// 创建临时存储的字节存储区
		bytes := make([]byte, 1024)
		read, err := connection.Read(bytes)
		if err != nil {
			fmt.Println("Read from client error:", err)
			fmt.Println("Connection with ", connection.RemoteAddr().String(), " down")
			break
		}
		fmt.Println(string(bytes[:read]))
	}
}
