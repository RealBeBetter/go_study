/**
 * @author Real
 * @since 2023/10/28 16:42
 */
package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// 直接运行 engine 服务
	err := engine.Run()
	if err != nil {
		return
	}

	// 指定端口运行 engine 服务
	// _ = engine.Run(":9999")

	e := gin.Default()
	e.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	// 指定端口运行 engine 服务
	_ = e.Run(":9999")
}
