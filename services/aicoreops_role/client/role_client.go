// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_role.proto

package client

import (
	"context"

	"aicoreops_role/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRoleRequest  = types.CreateRoleRequest
	CreateRoleResponse = types.CreateRoleResponse
	DeleteRoleRequest  = types.DeleteRoleRequest
	DeleteRoleResponse = types.DeleteRoleResponse
	GetRoleRequest     = types.GetRoleRequest
	GetRoleResponse    = types.GetRoleResponse
	ListRolesData      = types.ListRolesData
	ListRolesRequest   = types.ListRolesRequest
	ListRolesResponse  = types.ListRolesResponse
	Role               = types.Role
	UpdateRoleRequest  = types.UpdateRoleRequest
	UpdateRoleResponse = types.UpdateRoleResponse
	AssignPermissionsRequest = types.AssignPermissionsRequest
	AssignPermissionsResponse = types.AssignPermissionsResponse

	RoleService interface {
		// 创建角色
		CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
		// 获取角色详情
		GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error)
		// 更新角色
		UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
		// 删除角色
		DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error)
		// 列出角色
		ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
		// 分配权限
		AssignPermissions(ctx context.Context, in *AssignPermissionsRequest, opts ...grpc.CallOption) (*AssignPermissionsResponse, error)
	}

	defaultRoleService struct {
		cli zrpc.Client
	}
)

func NewRoleService(cli zrpc.Client) RoleService {
	return &defaultRoleService{
		cli: cli,
	}
}

// 创建角色
func (m *defaultRoleService) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	client := types.NewRoleServiceClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

// 获取角色详情
func (m *defaultRoleService) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error) {
	client := types.NewRoleServiceClient(m.cli.Conn())
	return client.GetRole(ctx, in, opts...)
}

// 更新角色
func (m *defaultRoleService) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	client := types.NewRoleServiceClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

// 删除角色
func (m *defaultRoleService) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error) {
	client := types.NewRoleServiceClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

// 列出角色
func (m *defaultRoleService) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	client := types.NewRoleServiceClient(m.cli.Conn())
	return client.ListRoles(ctx, in, opts...)
}

func (m *defaultRoleService) AssignPermissions(ctx context.Context, in *AssignPermissionsRequest, opts ...grpc.CallOption) (*AssignPermissionsResponse, error) {
	client := types.NewRoleServiceClient(m.cli.Conn())
	return client.AssignPermissions(ctx, in, opts...)
}