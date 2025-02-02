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
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/middleware"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	group := r.Group("/api")

	// 初始化各个 handler
	user := NewUserHandler(serverCtx)
	api := NewApiHandler(serverCtx)
	role := NewRoleHandler(serverCtx)
	menu := NewMenuHandler(serverCtx)
	ai := NewAiHandler(serverCtx)
	// 初始化中间件
	authMiddleware := middleware.NewAuthMiddleware(serverCtx.Config.JWT.Secret, serverCtx.RDB)
	casbinMiddleware := middleware.NewCasbinMiddleware(serverCtx.Enforcer)

	// 公开接口，不需要认证
	group.Post("/user/login", user.Login)
	group.Post("/user/create", user.CreateUser)

	// 需要认证和权限验证的接口组
	authGroup := group.Group("")
	authGroup.Use(authMiddleware.Handle, casbinMiddleware.Handle)

	// 用户相关接口
	authGroup.Post("/user/logout", user.Logout)
	authGroup.Get("/user/info", user.GetUserInfo)
	authGroup.Get("/user/get", user.GetUser)
	authGroup.Post("/user/update", user.UpdateUser)
	authGroup.Delete("/user/delete", user.DeleteUser)
	authGroup.Get("/user/list", user.ListUsers)
	authGroup.Get("/user/refresh", user.RefreshToken)
	authGroup.Get("/user/codes", user.GetAccessCodes)
	// API相关接口
	apiGroup := group.Group("")
	apiGroup.Use(authMiddleware.Handle, casbinMiddleware.Handle)
	apiGroup.Post("/api/create", api.CreateApi)
	apiGroup.Get("/api/get", api.GetApi)
	apiGroup.Post("/api/update", api.UpdateApi)
	apiGroup.Delete("/api/delete", api.DeleteApi)
	apiGroup.Get("/api/list", api.ListApis)

	// 角色相关接口
	roleGroup := group.Group("")
	roleGroup.Use(authMiddleware.Handle, casbinMiddleware.Handle)
	roleGroup.Post("/role/create", role.CreateRole)
	roleGroup.Get("/role/get", role.GetRole)
	roleGroup.Post("/role/update", role.UpdateRole)
	roleGroup.Delete("/role/delete", role.DeleteRole)
	roleGroup.Get("/role/list", role.ListRoles)
	roleGroup.Post("/role/assign_permissions", role.AssignPermissions)
	roleGroup.Post("/role/assign_role_to_user", role.AssignRoleToUser)
	roleGroup.Post("/role/remove_role_from_user", role.RemoveRoleFromUser)
	roleGroup.Post("/role/remove_permissions", role.RemoveUserPermissions)

	// 菜单相关接口
	menuGroup := group.Group("")
	menuGroup.Use(authMiddleware.Handle, casbinMiddleware.Handle)
	menuGroup.Post("/menu/create", menu.CreateMenu)
	menuGroup.Get("/menu/get", menu.GetMenu)
	menuGroup.Post("/menu/update", menu.UpdateMenu)
	menuGroup.Delete("/menu/delete", menu.DeleteMenu)
	menuGroup.Get("/menu/list", menu.ListMenus)

	// AI相关接口
	aiGroup := group.Group("")
	aiGroup.Use(authMiddleware.Handle) // , casbinMiddleware.Handle)
	aiGroup.Get("/ai/list", ai.GetChatList)
	aiGroup.Get("/ai/chat", ai.GetChatHistory)
	aiGroup.Post("/ai/upload", ai.UploadDocument)
	aiGroup.Get("/ai/ask", ai.AskQuestion)
	aiGroup.Post("/ai/newChat", ai.NewChat)
}
