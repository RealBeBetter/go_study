package gee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const ContentType = "Content-Type"

// H 构造 JSON 数据的别名
type H map[string]interface{}

type Context struct {
	// 原始对象
	Writer  http.ResponseWriter
	Request *http.Request

	// 请求信息
	Path   string
	Method string
	Params map[string]string

	// 返回信息
	StatusCode int

	// 中间件
	handlers []HandlerFunc
	index    int

	// 引擎指针
	engine *Engine
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		Params:  make(map[string]string),
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++

	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// ---------- 获取请求值 ----------

func (c *Context) GetPostFormVal(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) GetQueryVal(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) GetParamVal(key string) string {
	value := c.Params[key]
	return value
}

// ---------- 设置返回值 ----------

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Request.Header.Set(key, value)
}

// ---------- 设置返回值类型 ----------

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader(ContentType, "text/plain")
	c.SetStatusCode(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		log.Printf("Response String Error: %v", err)
	}
}

func (c *Context) JSON(code int, data interface{}) {
	c.SetHeader(ContentType, "application/json")
	c.SetStatusCode(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		log.Printf("Response JSON Error: %v", err)
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader(ContentType, "text/html")
	c.SetStatusCode(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		log.Printf("Response HTML Error: %v", err)
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatusCode(code)
	if _, err := c.Writer.Write(data); err != nil {
		log.Printf("Response Data Error: %v", err)
	}
}
