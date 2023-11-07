/**
 * @author wei.song
 * @since 2023/11/7 16:51
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}

// Engine is the uni handler for all requests
type Engine struct{}

// ServeHTTP implements the http.Handler interface
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}

}
