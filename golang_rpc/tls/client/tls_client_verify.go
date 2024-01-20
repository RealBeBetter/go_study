/**
 * @author Real
 * @since 2023/11/25 11:06
 */
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/rpc"
)

func main() {
	certPool := x509.NewCertPool()

	certBytes, err := ioutil.ReadFile("C:\\Users\\Real\\GoProjects\\go_study\\golang_rpc\\tls\\server\\server.crt")
	if err != nil {
		log.Fatalf("Failed to read server.crt: %v", err)
	}

	certPool.AppendCertsFromPEM(certBytes)

	config := &tls.Config{
		RootCAs: certPool,
	}

	connection, err := tls.Dial("tcp", "localhost:1234", config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer connection.Close()
	client := rpc.NewClient(connection)

	var result Result
	if err := client.Call("Calc.Square", 12, &result); err != nil {
		log.Fatalf("Failed to call Calc.Square: %v", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
