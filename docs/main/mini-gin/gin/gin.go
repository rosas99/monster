package gin

import (
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine struct
type Engine struct {
	middlewares []HandlerFunc
	routes      map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{
		routes:      make(map[string]HandlerFunc),
		middlewares: []HandlerFunc{},
	}
}

// Use adds middleware to the engine
func (e *Engine) Use(middlewares ...HandlerFunc) {
	e.middlewares = append(e.middlewares, middlewares...)
}

// AddRoute adds a route to the engine
func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.routes[key] = handler
}

// Run defines the method to start an HTTP server
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP implements the http.Handler interface
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	c.handlers = e.middlewares
	if handler, ok := e.routes[strings.ToUpper(req.Method)+"-"+req.URL.Path]; ok {
		c.handlers = append(c.handlers, handler)
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND")
		})
	}
	c.Next()
}
