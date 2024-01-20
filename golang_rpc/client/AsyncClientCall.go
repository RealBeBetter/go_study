/**
 * @author Real
 * @since 2023/11/22 23:34
 */
package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	var result Result
	asyncCall := client.Go("Calc.Square", 12, &result, nil)
	log.Printf("before call:%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("after call:%d^2 = %d", result.Num, result.Ans)
}
