/**
 * @author Real
 * @since 2023/11/25 10:52
 */
package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Calc int

func (cal *Calc) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	_ = rpc.Register(new(Calc))

	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Server RPC server on port %d", 1234)

	for {
		connection, _ := listener.Accept()

		defer connection.Close()

		go rpc.ServeConn(connection)
	}
}

// Context from Function C:/Users/Real/GoProjects/go_study/golang_rpc/server/tls_server.go:main.*Calc.Square
