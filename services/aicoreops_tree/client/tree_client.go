package client

import (
	"aicoreops_tree/types"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateTreeNodeRequest       = types.CreateTreeNodeRequest
	CreateTreeNodeResponse      = types.CreateTreeNodeResponse
	DeleteTreeNodeRequest       = types.DeleteTreeNodeRequest
	DeleteTreeNodeResponse      = types.DeleteTreeNodeResponse
	GetChildrenTreeNodeRequest  = types.GetChildrenTreeNodeRequest
	GetChildrenTreeNodeResponse = types.GetChildrenTreeNodeResponse
	GetTopTreeNodeRequest       = types.GetTopTreeNodeRequest
	GetTopTreeNodeResponse      = types.GetTopTreeNodeResponse
	ListLeafTreeNodeRequest     = types.ListLeafTreeNodeRequest
	ListLeafTreeNodeResponse    = types.ListLeafTreeNodeResponse
	ListTreeNodeRequest         = types.ListTreeNodeRequest
	ListTreeNodeResponse        = types.ListTreeNodeResponse
	ResourceTree                = types.ResourceTree
	SelectTreeNodeRequest       = types.SelectTreeNodeRequest
	SelectTreeNodeResponse      = types.SelectTreeNodeResponse
	SyncCMDBRequest             = types.SyncCMDBRequest
	SyncCMDBResponse            = types.SyncCMDBResponse
	UpdateTreeNodeRequest       = types.UpdateTreeNodeRequest
	UpdateTreeNodeResponse      = types.UpdateTreeNodeResponse

	ResourceTreeService interface {
		// 获取树节点列表
		ListTreeNode(ctx context.Context, in *ListTreeNodeRequest, opts ...grpc.CallOption) (*ListTreeNodeResponse, error)
		// 选择树节点
		SelectTreeNode(ctx context.Context, in *SelectTreeNodeRequest, opts ...grpc.CallOption) (*SelectTreeNodeResponse, error)
		// 获取顶级树节点
		GetTopTreeNode(ctx context.Context, in *GetTopTreeNodeRequest, opts ...grpc.CallOption) (*GetTopTreeNodeResponse, error)
		// 获取叶子节点列表
		ListLeafTreeNode(ctx context.Context, in *ListLeafTreeNodeRequest, opts ...grpc.CallOption) (*ListLeafTreeNodeResponse, error)
		// 创建树节点
		CreateTreeNode(ctx context.Context, in *CreateTreeNodeRequest, opts ...grpc.CallOption) (*CreateTreeNodeResponse, error)
		// 删除树节点
		DeleteTreeNode(ctx context.Context, in *DeleteTreeNodeRequest, opts ...grpc.CallOption) (*DeleteTreeNodeResponse, error)
		// 获取子节点
		GetChildrenTreeNode(ctx context.Context, in *GetChildrenTreeNodeRequest, opts ...grpc.CallOption) (*GetChildrenTreeNodeResponse, error)
		// 更新树节点
		UpdateTreeNode(ctx context.Context, in *UpdateTreeNodeRequest, opts ...grpc.CallOption) (*UpdateTreeNodeResponse, error)
		// 同步CMDB资源
		SyncCMDB(ctx context.Context, in *SyncCMDBRequest, opts ...grpc.CallOption) (*SyncCMDBResponse, error)
	}

	defaultResourceTreeService struct {
		cli zrpc.Client
	}
)

func NewResourceTreeService(cli zrpc.Client) ResourceTreeService {
	return &defaultResourceTreeService{
		cli: cli,
	}
}

// 获取树节点列表
func (m *defaultResourceTreeService) ListTreeNode(ctx context.Context, in *ListTreeNodeRequest, opts ...grpc.CallOption) (*ListTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.ListTreeNode(ctx, in, opts...)
}

// 选择树节点
func (m *defaultResourceTreeService) SelectTreeNode(ctx context.Context, in *SelectTreeNodeRequest, opts ...grpc.CallOption) (*SelectTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.SelectTreeNode(ctx, in, opts...)
}

// 获取顶级树节点
func (m *defaultResourceTreeService) GetTopTreeNode(ctx context.Context, in *GetTopTreeNodeRequest, opts ...grpc.CallOption) (*GetTopTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.GetTopTreeNode(ctx, in, opts...)
}

// 获取叶子节点列表
func (m *defaultResourceTreeService) ListLeafTreeNode(ctx context.Context, in *ListLeafTreeNodeRequest, opts ...grpc.CallOption) (*ListLeafTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.ListLeafTreeNode(ctx, in, opts...)
}

// 创建树节点
func (m *defaultResourceTreeService) CreateTreeNode(ctx context.Context, in *CreateTreeNodeRequest, opts ...grpc.CallOption) (*CreateTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.CreateTreeNode(ctx, in, opts...)
}

// 删除树节点
func (m *defaultResourceTreeService) DeleteTreeNode(ctx context.Context, in *DeleteTreeNodeRequest, opts ...grpc.CallOption) (*DeleteTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.DeleteTreeNode(ctx, in, opts...)
}

// 获取子节点
func (m *defaultResourceTreeService) GetChildrenTreeNode(ctx context.Context, in *GetChildrenTreeNodeRequest, opts ...grpc.CallOption) (*GetChildrenTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.GetChildrenTreeNode(ctx, in, opts...)
}

// 更新树节点
func (m *defaultResourceTreeService) UpdateTreeNode(ctx context.Context, in *UpdateTreeNodeRequest, opts ...grpc.CallOption) (*UpdateTreeNodeResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.UpdateTreeNode(ctx, in, opts...)
}

// 同步CMDB资源
func (m *defaultResourceTreeService) SyncCMDB(ctx context.Context, in *SyncCMDBRequest, opts ...grpc.CallOption) (*SyncCMDBResponse, error) {
	client := types.NewResourceTreeServiceClient(m.cli.Conn())
	return client.SyncCMDB(ctx, in, opts...)
}
