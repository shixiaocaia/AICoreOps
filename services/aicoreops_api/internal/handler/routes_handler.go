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
