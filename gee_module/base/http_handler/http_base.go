/**
 * @author Real
 * @since 2023/11/2 22:18
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "index hanlder, URL.path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for index, _ := range headers {
		_, err := fmt.Fprintf(w, "header[%s] = %s\n", index, headers[index])
		if err != nil {
			return
		}
	}
}
