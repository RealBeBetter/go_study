/**
 * @author Real
 * @since 2023/10/29 12:24
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("------------- routerWithRouter --------------")
	routerWithoutParameter()
}

func routerWithoutParameter() {
	// 创建一个默认的路由引擎
	engine := gin.Default()

	// 注册一个 GET 请求，第一个参数是请求的路径，第二个参数是处理这个请求的函数
	// gin.Context 封装了 request 和 response
	// context.String() 第一个参数是状态码，第二个参数是返回的内容
	// 没有指定端口，默认监听 8080 端口；没有参数，默认监听 / 路径
	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world!")
	})

	engine.Run()
}
