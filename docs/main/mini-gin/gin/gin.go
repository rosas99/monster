package gin

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	middlewares []HandlerFunc
	routes      map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		routes:      make(map[string]HandlerFunc),
		middlewares: []HandlerFunc{},
	}
}

func (e *Engine) Use(middlewares ...HandlerFunc) {
	e.middlewares = append(e.middlewares, middlewares...)
}

func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.routes[key] = handler
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

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
