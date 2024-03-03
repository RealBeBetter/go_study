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
	// server/simple_goroutine.go

	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	certPool := x509.NewCertPool()
	certBytes, _ := ioutil.ReadFile("../client/client.crt")
	certPool.AppendCertsFromPEM(certBytes)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	fmt.Println(config)
}
