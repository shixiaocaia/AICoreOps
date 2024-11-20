package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

type Routers struct {
	server      *rest.Server
	middlewares []rest.Middleware
	prefix      string
}

func NewRouters(server *rest.Server, prefix string) *Routers {
	return &Routers{
		server: server,
		prefix: prefix,
	}
}

// addRoute 内部通用路由添加方法
func (r *Routers) addRoute(method, path string, handler http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middlewares,
			rest.Route{
				Method:  method,
				Path:    path,
				Handler: handler,
			},
		),
		rest.WithPrefix(r.prefix),
	)
}

// Get 添加GET路由
func (r *Routers) Get(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodGet, path, handler)
}

// Post 添加POST路由
func (r *Routers) Post(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodPost, path, handler)
}

// Delete 添加DELETE路由
func (r *Routers) Delete(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodDelete, path, handler)
}

// Put 添加PUT路由
func (r *Routers) Put(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodPut, path, handler)
}

// Group 创建路由组
func (r *Routers) Group() *Routers {
	return &Routers{
		server: r.server,
		prefix: r.prefix,
	}
}

// Use 添加中间件
func (r *Routers) Use(middleware ...rest.Middleware) {
	r.middlewares = append(r.middlewares, middleware...)
}
