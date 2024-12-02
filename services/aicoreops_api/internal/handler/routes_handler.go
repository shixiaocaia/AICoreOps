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
 * File: routes_handler.go
 */

package handler

import (
	"aicoreops_api/internal/middleware"
	"aicoreops_api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	userGroup := r.Group("/api")

	user := NewUserHandler(serverCtx)
	authMiddleware := middleware.NewAuthMiddleware(serverCtx.Config.JWT.Secret, serverCtx.RDB)
	casbinMiddleware := middleware.NewCasbinMiddleware(serverCtx.Enforcer)

	// login 接口不需要认证
	userGroup.Post("/user/login", user.Login)
	userGroup.Post("/user/create", user.CreateUser)

	// 其他接口需要认证和权限验证
	authGroup := userGroup.Group("")
	authGroup.Use(authMiddleware.Handle, casbinMiddleware.Handle)

	authGroup.Post("/user/logout", user.Logout)
	authGroup.Get("/user/get", user.GetUser)
	authGroup.Post("/user/update", user.UpdateUser)
	authGroup.Delete("/user/delete", user.DeleteUser)
	authGroup.Get("/user/list", user.ListUsers)
}
