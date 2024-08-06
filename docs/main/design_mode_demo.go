package main

import (
	"fmt"
	"net/http"
)

// Engine 用于管理中间件和路由
type Engine struct {
	middlewares []Middleware // 中间件列表
}

// Middleware 函数类型，用于处理HTTP请求和响应
type Middleware func(*Context)

// Context 封装了http.ResponseWriter和*http.Request
type Context struct {
	http.ResponseWriter
	*http.Request
	index  int // 当前中间件索引
	Engine *Engine
}

// NewEngine 创建并返回一个新的Engine实例
func NewEngine() *Engine {
	return &Engine{
		middlewares: []Middleware{},
	}
}

// Use 添加中间件到Engine
func (e *Engine) Use(middlewares ...Middleware) {
	e.middlewares = append(e.middlewares, middlewares...)
}

// ServeHTTP 实现http.Handler接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		ResponseWriter: w,
		Request:        r,
		index:          len(e.middlewares),
		Engine:         e,
	}

	c.next()
}

// next 调用链中的下一个中间件
func (c *Context) next() {
	c.index--
	if c.index < 0 {
		// 没有更多的中间件，调用最终的处理器
		http.NotFound(c.ResponseWriter, c.Request)
		return
	}
	middleware := c.Engine.middlewares[c.index]
	middleware(c)
}

// LoggerMiddleware 示例中间件，记录日志
func LoggerMiddleware(c *Context) {
	fmt.Printf("%s %s\n", c.Request.Method, c.Request.URL.Path)
	c.next() // 调用下一个中间件
}

// AuthMiddleware 示例中间件，进行认证
func AuthMiddleware(c *Context) {
	// 这里使用简单的认证逻辑
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader != "Bearer SecretToken" {
		http.Error(c.ResponseWriter, "Unauthorized", http.StatusUnauthorized)
		return
	}
	c.next() // 调用下一个中间件
}

func main() {
	router := NewEngine()

	// 使用中间件
	router.Use(LoggerMiddleware, AuthMiddleware)

	// 注册路由和处理器
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! This is the final handler.")
	}))

	// 启动服务器
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// Handle 方法用于注册路由和处理器
func (e *Engine) Handle(pattern string, handler http.Handler) {
	// 这里简化实现，直接使用 http.ServeMux 来处理路由
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
