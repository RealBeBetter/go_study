package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=Real
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetQueryVal("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/Real
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetParamVal("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.GetPostFormVal("username"),
				"password": c.GetPostFormVal("password"),
			})
		})

	}

	r.Run(":9999")
}
