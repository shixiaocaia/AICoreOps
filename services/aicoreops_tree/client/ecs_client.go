package client

import (
	"aicoreops_tree/types"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchOperateEcsRequest   = types.BatchOperateEcsRequest
	BatchOperateEcsResponse  = types.BatchOperateEcsResponse
	BindEcsRequest           = types.BindEcsRequest
	BindEcsResponse          = types.BindEcsResponse
	EcsInstance              = types.EcsInstance
	GetEcsListRequest        = types.GetEcsListRequest
	GetEcsListResponse       = types.GetEcsListResponse
	GetEcsUnbindListRequest  = types.GetEcsUnbindListRequest
	GetEcsUnbindListResponse = types.GetEcsUnbindListResponse
	UnBindEcsRequest         = types.UnBindEcsRequest
	UnBindEcsResponse        = types.UnBindEcsResponse

	EcsService interface {
		// 获取未绑定的ECS列表
		GetEcsUnbindList(ctx context.Context, in *GetEcsUnbindListRequest, opts ...grpc.CallOption) (*GetEcsUnbindListResponse, error)
		// 获取已绑定的ECS列表
		GetEcsList(ctx context.Context, in *GetEcsListRequest, opts ...grpc.CallOption) (*GetEcsListResponse, error)
		// 绑定ECS到资源树节点
		BindEcs(ctx context.Context, in *BindEcsRequest, opts ...grpc.CallOption) (*BindEcsResponse, error)
		// 从资源树节点解绑ECS
		UnBindEcs(ctx context.Context, in *UnBindEcsRequest, opts ...grpc.CallOption) (*UnBindEcsResponse, error)
		// 批量操作ECS实例
		BatchOperateEcs(ctx context.Context, in *BatchOperateEcsRequest, opts ...grpc.CallOption) (*BatchOperateEcsResponse, error)
	}

	defaultEcsService struct {
		cli zrpc.Client
	}
)

func NewEcsService(cli zrpc.Client) EcsService {
	return &defaultEcsService{
		cli: cli,
	}
}

// 获取未绑定的ECS列表
func (m *defaultEcsService) GetEcsUnbindList(ctx context.Context, in *GetEcsUnbindListRequest, opts ...grpc.CallOption) (*GetEcsUnbindListResponse, error) {
	client := types.NewEcsServiceClient(m.cli.Conn())
	return client.GetEcsUnbindList(ctx, in, opts...)
}

// 获取已绑定的ECS列表
func (m *defaultEcsService) GetEcsList(ctx context.Context, in *GetEcsListRequest, opts ...grpc.CallOption) (*GetEcsListResponse, error) {
	client := types.NewEcsServiceClient(m.cli.Conn())
	return client.GetEcsList(ctx, in, opts...)
}

// 绑定ECS到资源树节点
func (m *defaultEcsService) BindEcs(ctx context.Context, in *BindEcsRequest, opts ...grpc.CallOption) (*BindEcsResponse, error) {
	client := types.NewEcsServiceClient(m.cli.Conn())
	return client.BindEcs(ctx, in, opts...)
}

// 从资源树节点解绑ECS
func (m *defaultEcsService) UnBindEcs(ctx context.Context, in *UnBindEcsRequest, opts ...grpc.CallOption) (*UnBindEcsResponse, error) {
	client := types.NewEcsServiceClient(m.cli.Conn())
	return client.UnBindEcs(ctx, in, opts...)
}

// 批量操作ECS实例
func (m *defaultEcsService) BatchOperateEcs(ctx context.Context, in *BatchOperateEcsRequest, opts ...grpc.CallOption) (*BatchOperateEcsResponse, error) {
	client := types.NewEcsServiceClient(m.cli.Conn())
	return client.BatchOperateEcs(ctx, in, opts...)
}
