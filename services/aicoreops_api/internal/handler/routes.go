/*
 * Copyright 2024 Bamboo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * File: routes.go
 */

package handler

import (
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/rest"
)

type Routers struct {
	server      *rest.Server
	middlewares []rest.Middleware
	prefix      string
}

func NewRouters(server *rest.Server) *Routers {
	return &Routers{
		server: server,
	}
}

// addRoute 内部通用路由添加方法
func (r *Routers) addRoute(method, routePath string, handler http.HandlerFunc) *Routers {
	fullPath := path.Join(r.prefix, routePath)
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middlewares,
			rest.Route{
				Method:  method,
				Path:    fullPath,
				Handler: handler,
			},
		),
	)
	return r
}

// Get 添加GET路由
func (r *Routers) Get(path string, handler http.HandlerFunc) *Routers {
	return r.addRoute(http.MethodGet, path, handler)
}

// Post 添加POST路由
func (r *Routers) Post(path string, handler http.HandlerFunc) *Routers {
	return r.addRoute(http.MethodPost, path, handler)
}

// Delete 添加DELETE路由
func (r *Routers) Delete(path string, handler http.HandlerFunc) *Routers {
	return r.addRoute(http.MethodDelete, path, handler)
}

// Put 添加PUT路由
func (r *Routers) Put(path string, handler http.HandlerFunc) *Routers {
	return r.addRoute(http.MethodPut, path, handler)
}

// Group 创建路由组
func (r *Routers) Group(prefix ...string) *Routers {
	group := &Routers{
		server: r.server,
	}
	if len(prefix) > 0 {
		group.prefix = path.Join(r.prefix, prefix[0])
	}
	return group
}

// Use 添加中间件
func (r *Routers) Use(middleware ...rest.Middleware) *Routers {
	r.middlewares = append(r.middlewares, middleware...)
	return r
}
