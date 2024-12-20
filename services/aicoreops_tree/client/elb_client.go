package client

import (
	"aicoreops_tree/types"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchOperateElbRequest   = types.BatchOperateElbRequest
	BatchOperateElbResponse  = types.BatchOperateElbResponse
	BindElbRequest           = types.BindElbRequest
	BindElbResponse          = types.BindElbResponse
	ElbInstance              = types.ElbInstance
	GetElbListRequest        = types.GetElbListRequest
	GetElbListResponse       = types.GetElbListResponse
	GetElbUnbindListRequest  = types.GetElbUnbindListRequest
	GetElbUnbindListResponse = types.GetElbUnbindListResponse
	UnBindElbRequest         = types.UnBindElbRequest
	UnBindElbResponse        = types.UnBindElbResponse

	ElbService interface {
		// 获取未绑定的ELB列表
		GetElbUnbindList(ctx context.Context, in *GetElbUnbindListRequest, opts ...grpc.CallOption) (*GetElbUnbindListResponse, error)
		// 获取已绑定的ELB列表
		GetElbList(ctx context.Context, in *GetElbListRequest, opts ...grpc.CallOption) (*GetElbListResponse, error)
		// 绑定ELB到资源树节点
		BindElb(ctx context.Context, in *BindElbRequest, opts ...grpc.CallOption) (*BindElbResponse, error)
		// 从资源树节点解绑ELB
		UnBindElb(ctx context.Context, in *UnBindElbRequest, opts ...grpc.CallOption) (*UnBindElbResponse, error)
		// 批量操作ELB实例
		BatchOperateElb(ctx context.Context, in *BatchOperateElbRequest, opts ...grpc.CallOption) (*BatchOperateElbResponse, error)
	}

	defaultElbService struct {
		cli zrpc.Client
	}
)

func NewElbService(cli zrpc.Client) ElbService {
	return &defaultElbService{
		cli: cli,
	}
}

// 获取未绑定的ELB列表
func (m *defaultElbService) GetElbUnbindList(ctx context.Context, in *GetElbUnbindListRequest, opts ...grpc.CallOption) (*GetElbUnbindListResponse, error) {
	client := types.NewElbServiceClient(m.cli.Conn())
	return client.GetElbUnbindList(ctx, in, opts...)
}

// 获取已绑定的ELB列表
func (m *defaultElbService) GetElbList(ctx context.Context, in *GetElbListRequest, opts ...grpc.CallOption) (*GetElbListResponse, error) {
	client := types.NewElbServiceClient(m.cli.Conn())
	return client.GetElbList(ctx, in, opts...)
}

// 绑定ELB到资源树节点
func (m *defaultElbService) BindElb(ctx context.Context, in *BindElbRequest, opts ...grpc.CallOption) (*BindElbResponse, error) {
	client := types.NewElbServiceClient(m.cli.Conn())
	return client.BindElb(ctx, in, opts...)
}

// 从资源树节点解绑ELB
func (m *defaultElbService) UnBindElb(ctx context.Context, in *UnBindElbRequest, opts ...grpc.CallOption) (*UnBindElbResponse, error) {
	client := types.NewElbServiceClient(m.cli.Conn())
	return client.UnBindElb(ctx, in, opts...)
}

// 批量操作ELB实例
func (m *defaultElbService) BatchOperateElb(ctx context.Context, in *BatchOperateElbRequest, opts ...grpc.CallOption) (*BatchOperateElbResponse, error) {
	client := types.NewElbServiceClient(m.cli.Conn())
	return client.BatchOperateElb(ctx, in, opts...)
}
