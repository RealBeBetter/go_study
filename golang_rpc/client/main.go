/**
 * @author Real
 * @since 2023/11/21 22:17
 */
package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	var result Result
	if err := client.Call("Calc.Square", 12, &result); err != nil {
		log.Fatalf("Failed to call to Calc.Square, error = %v\n", err)
	}

	log.Printf("%d^2 = %d\n", result.Num, result.Ans)
}
