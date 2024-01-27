package gee_context

import (
	"gee/gee"
	"net/http"
	"testing"
)

func TestContext(t *testing.T) {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=Real
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetQueryVal("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.GetPostFormVal("username"),
			"password": c.GetPostFormVal("password"),
		})
	})

	if err := r.Run(":9999"); err != nil {
		t.Fatalf("Error, exit. error: %v", err)
	}
}
