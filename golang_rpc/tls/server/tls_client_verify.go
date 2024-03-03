/**
 * @author Real
 * @since 2023/11/26 1:00
 */
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func main() {
	// client/simple_goroutine.go

	cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
	certPool := x509.NewCertPool()
	certBytes, _ := ioutil.ReadFile("C:\\Users\\Real\\GoProjects\\go_study\\golang_rpc\\tls\\server\\server.crt")
	certPool.AppendCertsFromPEM(certBytes)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	fmt.Println(config)
}
