// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_tree.proto

package server

import (
	"context"
	"aicoreops_tree/internal/logic"
	"aicoreops_tree/internal/model"
	"aicoreops_tree/internal/svc"
	"aicoreops_tree/types"
)

type ResourceTreeServiceServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedResourceTreeServiceServer
}

func NewResourceTreeServiceServer(svcCtx *svc.ServiceContext) *ResourceTreeServiceServer {
	return &ResourceTreeServiceServer{
		svcCtx: svcCtx,
	}
}

// ListTreeNode 查询树节点列表
func (s *ResourceTreeServiceServer) ListTreeNode(ctx context.Context, req *types.ListTreeNodeRequest) (*types.ListTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	nodes, total, err := l.ListTreeNode(ctx, req.PageSize, req.PageNum, req.CmdbType, req.Status, req.Keyword)
	if err != nil {
		return nil, err
	}
	return &types.ListTreeNodeResponse{
		Total: total,
		Nodes: convertTreeNodes(nodes),
	}, nil
}

// SelectTreeNode 查询单个树节点
func (s *ResourceTreeServiceServer) SelectTreeNode(ctx context.Context, req *types.SelectTreeNodeRequest) (*types.SelectTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	node, err := l.SelectTreeNode(ctx, req.Id, req.CmdbId)
	if err != nil {
		return nil, err
	}
	return &types.SelectTreeNodeResponse{
		Node: convertTreeNode(node),
		Exists: node != nil,
	}, nil
}

// GetTopTreeNode 查询顶级树节点
func (s *ResourceTreeServiceServer) GetTopTreeNode(ctx context.Context, req *types.GetTopTreeNodeRequest) (*types.GetTopTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	nodes, total, err := l.GetTopTreeNode(ctx, req.Limit, req.CmdbType, req.Status)
	if err != nil {
		return nil, err
	}
	return &types.GetTopTreeNodeResponse{
		Total: total,
		Nodes: convertTreeNodes(nodes),
	}, nil
}

// ListLeafTreeNode 查询叶子节点列表
func (s *ResourceTreeServiceServer) ListLeafTreeNode(ctx context.Context, req *types.ListLeafTreeNodeRequest) (*types.ListLeafTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	nodes, total, err := l.ListLeafTreeNode(ctx, req.PageSize, req.PageNum, req.CmdbType, req.Status)
	if err != nil {
		return nil, err
	}
	return &types.ListLeafTreeNodeResponse{
		Total: total,
		Nodes: convertTreeNodes(nodes),
	}, nil
}

// CreateTreeNode 创建树节点
func (s *ResourceTreeServiceServer) CreateTreeNode(ctx context.Context, req *types.CreateTreeNodeRequest) (*types.CreateTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	node := &model.TreeNode{
		Pid: int(req.Pid),
		Title: req.Title,
		Description: req.Description,
		IsLeaf: int(req.IsLeaf),
		CMDBID: req.CmdbId,
		CMDBType: req.CmdbType,
		CMDBAttrs: req.CmdbAttrs,
		Creator: req.Creator,
	}
	err := l.CreateTreeNode(ctx, node)
	if err != nil {
		return nil, err
	}
	return &types.CreateTreeNodeResponse{
		Success: true,
		Node: convertTreeNode(node),
	}, nil
}

// DeleteTreeNode 删除树节点
func (s *ResourceTreeServiceServer) DeleteTreeNode(ctx context.Context, req *types.DeleteTreeNodeRequest) (*types.DeleteTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	err := l.DeleteTreeNode(ctx, req.Id, req.Operator)
	if err != nil {
		return nil, err
	}
	return &types.DeleteTreeNodeResponse{
		Success: true,
	}, nil
}

// GetChildrenTreeNode 获取子节点列表
func (s *ResourceTreeServiceServer) GetChildrenTreeNode(ctx context.Context, req *types.GetChildrenTreeNodeRequest) (*types.GetChildrenTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	nodes, total, err := l.GetChildrenTreeNode(ctx, req.Pid, req.PageSize, req.PageNum, req.CmdbType, req.Status)
	if err != nil {
		return nil, err
	}
	return &types.GetChildrenTreeNodeResponse{
		Total: total,
		Nodes: convertTreeNodes(nodes),
	}, nil
}

// UpdateTreeNode 更新树节点信息
func (s *ResourceTreeServiceServer) UpdateTreeNode(ctx context.Context, req *types.UpdateTreeNodeRequest) (*types.UpdateTreeNodeResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	node := &model.TreeNode{
		ID: req.Id,
		Title: req.Title,
		Description: req.Description,
		IsLeaf: int(req.IsLeaf),
		CMDBAttrs: req.CmdbAttrs,
		Updater: req.Updater,
		Status: req.Status,
	}
	err := l.UpdateTreeNode(ctx, node)
	if err != nil {
		return nil, err
	}
	return &types.UpdateTreeNodeResponse{
		Success: true,
		Node: convertTreeNode(node),
	}, nil
}

// SyncCMDB 同步CMDB资源数据到树节点
func (s *ResourceTreeServiceServer) SyncCMDB(ctx context.Context, req *types.SyncCMDBRequest) (*types.SyncCMDBResponse, error) {
	l := logic.NewTreeLogic(ctx, s.svcCtx)
	node := &model.TreeNode{
		CMDBType: req.CmdbType,
		Updater: req.Operator,
	}
	err := l.SyncCMDB(ctx, node)
	if err != nil {
		return nil, err
	}
	return &types.SyncCMDBResponse{
		Success: true,
	}, nil
}

// 工具函数:将model.TreeNode转换为types.ResourceTree
func convertTreeNode(node *model.TreeNode) *types.ResourceTree {
	if node == nil {
		return nil
	}
	return &types.ResourceTree{
		Id: node.ID,
		CreateTime: node.CreateTime,
		UpdateTime: node.UpdateTime,
		IsDeleted: int32(node.IsDeleted),
		Title: node.Title,
		Pid: int32(node.Pid),
		Level: int32(node.Level),
		IsLeaf: int32(node.IsLeaf),
		Description: node.Description,
		CmdbId: node.CMDBID,
		CmdbType: node.CMDBType,
		CmdbAttrs: node.CMDBAttrs,
		Creator: node.Creator,
		Updater: node.Updater,
		Status: node.Status,
	}
}

// 工具函数:将model.TreeNode切片转换为types.ResourceTree切片
func convertTreeNodes(nodes []*model.TreeNode) []*types.ResourceTree {
	if nodes == nil {
		return nil
	}
	result := make([]*types.ResourceTree, len(nodes))
	for i, node := range nodes {
		result[i] = convertTreeNode(node)
	}
	return result
}