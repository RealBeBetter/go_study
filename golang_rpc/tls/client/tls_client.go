/**
 * @author Real
 * @since 2023/11/25 10:58
 */
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "localhost:1234", config)
	if err != nil {
		fmt.Printf("Failed to dial. error: %v\n", err)
		return
	}

	defer func(conn *tls.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Failed to close connection. ", err)
		}
	}(conn)
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Calc.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Calc.Square. ", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
