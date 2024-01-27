package gee

import (
	"log"
	"net/http"
)

type Router struct {
	Handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("[Route]: %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.Handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.Handlers[key]; ok {
		handler(c)
		return
	}

	// not found, return 404
	c.String(http.StatusNotFound, "[Route]: 404 NOT FOUND: %s\n", c.Path)
}
