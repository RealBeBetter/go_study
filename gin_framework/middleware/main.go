/**
 * @author wei.song
 * @since 2023/10/30 18:51
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	fmt.Println("------------- middlewareGlobal -------------")
	// middlewareGlobal()

	fmt.Println("------------- middlewareSingle -------------")
	// middlewareSingle()

	fmt.Println("------------- middlewareGroup -------------")
	middlewareGroup()
}

func middlewareGlobal() {
	engine := gin.Default()
	// 作用于全局
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 作用于全局
	engine.Use(Logger())

	engine.GET("/hello", func(context *gin.Context) {
		keys := context.Keys
		fmt.Println(keys)
		context.JSON(200, gin.H{
			"message": "hello world",
			"Test":    context.GetString("Test"),
		})
	})

	_ = engine.Run(":9999")
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		// 给Context实例设置一个值
		context.Set("Test", "123456")
		// 请求前
		context.Next()
		// 请求后
		latency := time.Since(now)
		log.Print(latency)
	}
}

func middlewareSingle() {
	engine := gin.Default()
	// 作用于单个路由
	engine.GET("/benchmark", Logger())

	_ = engine.Run(":9999")
}

func middlewareGroup() {
	engine := gin.Default()

	// 作用于某个组
	authorized := engine.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint())
		authorized.POST("/submit", submitEndpoint())
	}

	_ = engine.Run(":9999")
}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetBool("authenticated") {
			context.Next()
		} else {
			context.String(401, "unauthorized")
		}
	}
}

func loginEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, "login")
	}
}

func submitEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, "submit")
	}
}
