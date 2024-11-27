package handler

import (
	"aicoreops_api/internal/middleware"
	"aicoreops_api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	userGroup := r.Group("/api")
	user := NewUserHandler(serverCtx)
	userGroup.Post("/user/login", user.Login)
	userGroup.Post("/user/logout", user.Logout).Use(middleware.AuthMiddleware(serverCtx.Config.JWT.Secret))
	userGroup.Post("/user/create", user.CreateUser)
	userGroup.Get("/user/get", user.GetUser).Use(middleware.AuthMiddleware(serverCtx.Config.JWT.Secret))
	userGroup.Post("/user/update", user.UpdateUser).Use(middleware.AuthMiddleware(serverCtx.Config.JWT.Secret))
	userGroup.Delete("/user/delete", user.DeleteUser).Use(middleware.AuthMiddleware(serverCtx.Config.JWT.Secret))
	userGroup.Get("/user/list", user.ListUsers).Use(middleware.AuthMiddleware(serverCtx.Config.JWT.Secret))
}
