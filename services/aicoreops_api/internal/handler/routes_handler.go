package handler

import (
	"aicoreops_api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	userGroup := r.Group("/api")
	user := NewUserHandler(serverCtx)
	userGroup.Post("/user/login", user.Login)
	userGroup.Post("/user/logout", user.Logout)
	userGroup.Post("/user/create", user.Register)
	userGroup.Get("/user/get", user.GetUser)
	userGroup.Post("/user/update", user.UpdateUser)
	userGroup.Delete("/user/delete", user.DeleteUser)
	userGroup.Get("/user/list", user.ListUsers)
}
