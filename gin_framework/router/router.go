/**
 * @author Real
 * @since 2023/10/29 13:20
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("------------- routerWithRouter --------------")
	// routerWithoutParameter()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithPathParameter()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithQuery()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithPostForm()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithQueryAndPostForm()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithMap()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithRedirect()

	fmt.Println("------------- routerWithRouter --------------")
	// routerWithRedirectSelf()

	fmt.Println("------------- routerWithRouter --------------")
	routerWithGroup()
}

func routerWithoutParameter() {
	// 创建一个默认的路由引擎
	engine := gin.Default()

	// 注册一个 GET 请求，第一个参数是请求的路径，第二个参数是处理这个请求的函数
	// gin_framework.Context 封装了 request 和 response
	// context.String() 第一个参数是状态码，第二个参数是返回的内容
	// 没有指定端口，默认监听 8080 端口；没有参数，默认监听 / 路径
	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world!")
	})

	_ = engine.Run()
}

func routerWithPathParameter() {
	engine := gin.Default()
	// func 程序中的函数没有名称，称为匿名函数
	engine.GET("/hello/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(200, "hello %s", name)
	})

	_ = engine.Run()
}

// 匹配users?name=xxx&role=xxx，role可选
func routerWithQuery() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		name := context.Query("name")
		role := context.Query("role")
		context.String(200, "%s is a %s", name, role)
	})

	_ = engine.Run()
}

func routerWithPostForm() {
	engine := gin.Default()
	engine.POST("/login", func(context *gin.Context) {
		username := context.PostForm("username")
		// 优先以 postForm 的值为准，如果没有则使用默认值
		password := context.DefaultPostForm("password", "000000")
		context.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})

	_ = engine.Run()
}

func routerWithQueryAndPostForm() {
	engine := gin.Default()
	engine.POST("/pages", func(context *gin.Context) {
		pageNum := context.Query("pageNum")
		pageSize := context.DefaultQuery("pageSize", "10")
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "000000")

		context.JSON(http.StatusOK, gin.H{
			"pageNum":  pageNum,
			"pageSize": pageSize,
			"username": username,
			"password": password,
		})
	})

	_ = engine.Run(":9999")
}

func routerWithMap() {
	engine := gin.Default()
	engine.POST("/map", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		context.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	_ = engine.Run(":9999")
}

func routerWithRedirect() {
	engine := gin.Default()
	engine.GET("/redirectBaidu", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	_ = engine.Run(":9999")
}

func routerWithRedirectSelf() {
	engine := gin.Default()
	engine.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index")
	})

	engine.GET("/index", func(context *gin.Context) {
		context.Request.URL.Path = "/"
		context.String(http.StatusOK, "index")
	})

	_ = engine.Run(":9999")
}

// 分组路由，一般用于相同前缀的路径分组
func routerWithGroup() {
	defaultHandler := func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"path": context.FullPath(),
		})
	}

	engine := gin.Default()
	v1Group := engine.Group("/api/v1")
	{
		v1Group.GET("/hello", defaultHandler)
		v1Group.GET("/greet", defaultHandler)
	}

	v2Group := engine.Group("/api/v2")
	{
		v2Group.GET("/hello", defaultHandler)
		v2Group.GET("/greet", defaultHandler)
	}

	_ = engine.Run(":9999")
}
