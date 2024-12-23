package client

import (
	"aicoreops_tree/types"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchOperateRdsRequest   = types.BatchOperateRdsRequest
	BatchOperateRdsResponse  = types.BatchOperateRdsResponse
	BindRdsRequest           = types.BindRdsRequest
	BindRdsResponse          = types.BindRdsResponse
	GetRdsListRequest        = types.GetRdsListRequest
	GetRdsListResponse       = types.GetRdsListResponse
	GetRdsUnbindListRequest  = types.GetRdsUnbindListRequest
	GetRdsUnbindListResponse = types.GetRdsUnbindListResponse
	RdsInstance              = types.RdsInstance
	UnBindRdsRequest         = types.UnBindRdsRequest
	UnBindRdsResponse        = types.UnBindRdsResponse

	RdsService interface {
		// 获取未绑定的RDS列表
		GetRdsUnbindList(ctx context.Context, in *GetRdsUnbindListRequest, opts ...grpc.CallOption) (*GetRdsUnbindListResponse, error)
		// 获取已绑定的RDS列表
		GetRdsList(ctx context.Context, in *GetRdsListRequest, opts ...grpc.CallOption) (*GetRdsListResponse, error)
		// 绑定RDS到资源树节点
		BindRds(ctx context.Context, in *BindRdsRequest, opts ...grpc.CallOption) (*BindRdsResponse, error)
		// 从资源树节点解绑RDS
		UnBindRds(ctx context.Context, in *UnBindRdsRequest, opts ...grpc.CallOption) (*UnBindRdsResponse, error)
		// 批量操作RDS实例
		BatchOperateRds(ctx context.Context, in *BatchOperateRdsRequest, opts ...grpc.CallOption) (*BatchOperateRdsResponse, error)
	}

	defaultRdsService struct {
		cli zrpc.Client
	}
)

func NewRdsService(cli zrpc.Client) RdsService {
	return &defaultRdsService{
		cli: cli,
	}
}

// 获取未绑定的RDS列表
func (m *defaultRdsService) GetRdsUnbindList(ctx context.Context, in *GetRdsUnbindListRequest, opts ...grpc.CallOption) (*GetRdsUnbindListResponse, error) {
	client := types.NewRdsServiceClient(m.cli.Conn())
	return client.GetRdsUnbindList(ctx, in, opts...)
}

// 获取已绑定的RDS列表
func (m *defaultRdsService) GetRdsList(ctx context.Context, in *GetRdsListRequest, opts ...grpc.CallOption) (*GetRdsListResponse, error) {
	client := types.NewRdsServiceClient(m.cli.Conn())
	return client.GetRdsList(ctx, in, opts...)
}

// 绑定RDS到资源树节点
func (m *defaultRdsService) BindRds(ctx context.Context, in *BindRdsRequest, opts ...grpc.CallOption) (*BindRdsResponse, error) {
	client := types.NewRdsServiceClient(m.cli.Conn())
	return client.BindRds(ctx, in, opts...)
}

// 从资源树节点解绑RDS
func (m *defaultRdsService) UnBindRds(ctx context.Context, in *UnBindRdsRequest, opts ...grpc.CallOption) (*UnBindRdsResponse, error) {
	client := types.NewRdsServiceClient(m.cli.Conn())
	return client.UnBindRds(ctx, in, opts...)
}

// 批量操作RDS实例
func (m *defaultRdsService) BatchOperateRds(ctx context.Context, in *BatchOperateRdsRequest, opts ...grpc.CallOption) (*BatchOperateRdsResponse, error) {
	client := types.NewRdsServiceClient(m.cli.Conn())
	return client.BatchOperateRds(ctx, in, opts...)
}
