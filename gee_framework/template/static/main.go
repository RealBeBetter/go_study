package main

import "gee/gee"

func main() {
	r := gee.New()
	r.Static("/assets", "/usr/Real/blog/static")

	// 或相对路径 r.Static("/assets", "./static")
	r.Run(":9999")
}
