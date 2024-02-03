package gee

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	Roots    map[string]*Node       // 根节点，由方法确定
	Handlers map[string]HandlerFunc // 对应的 HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func NewRouter() *Router {
	return &Router{
		Roots:    make(map[string]*Node),
		Handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 只允许路由中有一个 * 通配符，解析出每一块 / 分割的值
func parsePattern(pattern string) []string {
	patternPaths := strings.Split(pattern, "/")

	parts := make([]string, 0, len(patternPaths))
	for _, part := range patternPaths {
		// 路径为空
		if len(part) == 0 {
			continue
		}

		parts = append(parts, part)
		if strings.HasPrefix(part, "*") {
			break
		}
	}

	return parts
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	method = strings.ToUpper(method)
	log.Printf("[Route]: %4s - %s", method, pattern)

	parts := parsePattern(pattern)

	key := method + "-" + pattern
	if _, ok := r.Roots[method]; !ok {
		// 没有该类型的方法，新建节点
		r.Roots[method] = &Node{}
	}

	// 添加路由与 HandlerFunc
	r.Roots[method].Insert(pattern, parts, 0)
	r.Handlers[key] = handler
}

func (r *Router) getRoute(method string, path string) (node *Node, params map[string]string) {
	method = strings.ToUpper(method)
	root, ok := r.Roots[method]
	if !ok {
		return nil, nil
	}

	searchParts := parsePattern(path)
	node = root.Search(searchParts, 0)
	if node == nil {
		return nil, nil
	}

	params = make(map[string]string)

	parts := parsePattern(node.Pattern)
	for index, part := range parts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = searchParts[index]
		}

		if strings.HasPrefix(part, "*") && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}

	return node, params
}

func (r *Router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node == nil {
		// not found, return 404
		c.String(http.StatusNotFound, "[Route]: 404 NOT FOUND: %s\n", c.Path)
	}

	key := c.Method + "-" + c.Path
	c.Params = params
	handlerFunc := r.Handlers[key]
	handlerFunc(c)
}
