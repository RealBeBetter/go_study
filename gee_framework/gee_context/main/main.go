package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(writer http.ResponseWriter, request *http.Request) {
	})
}
