/**
 * @author wei.song
 * @since 2023/10/30 17:20
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("------------- HTML Template --------------")
	htmlTemplate()
}

func htmlTemplate() {
	engine := gin.Default()
	engine.LoadHTMLGlob("/Users/*/GolandProjects/go_study/gin_framework/template/templates/*")

	stu1 := &student{Name: "Real", Age: 18}
	stu2 := &student{Name: "Alice", Age: 18}

	engine.GET("/arr", func(context *gin.Context) {
		context.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	_ = engine.Run(":9999")
}
