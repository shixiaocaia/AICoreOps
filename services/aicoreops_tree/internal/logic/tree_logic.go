package logic

import (
	"aicoreops_tree/internal/domain"
	"aicoreops_tree/internal/model"
	"aicoreops_tree/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	treeDomain *domain.TreeDomain
	logx.Logger
}

func NewTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeLogic {
	return &TreeLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		treeDomain: domain.NewTreeDomain(svcCtx.DB),
		Logger:     logx.WithContext(ctx),
	}
}

// ListTreeNode 查询树节点列表
func (l *TreeLogic) ListTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status, keyword string) ([]*model.TreeNode, int32, error) {
	return l.treeDomain.ListTreeNode(ctx, pageSize, pageNum, cmdbType, status, keyword)
}

// SelectTreeNode 查询单个树节点
func (l *TreeLogic) SelectTreeNode(ctx context.Context, id int64, cmdbID string) (*model.TreeNode, error) {
	return l.treeDomain.SelectTreeNode(ctx, id, cmdbID)
}

// GetTopTreeNode 查询顶级树节点
func (l *TreeLogic) GetTopTreeNode(ctx context.Context, limit int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	return l.treeDomain.GetTopTreeNode(ctx, limit, cmdbType, status)
}

// ListLeafTreeNode 查询叶子节点列表
func (l *TreeLogic) ListLeafTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	return l.treeDomain.ListLeafTreeNode(ctx, pageSize, pageNum, cmdbType, status)
}

// CreateTreeNode 创建树节点
func (l *TreeLogic) CreateTreeNode(ctx context.Context, node *model.TreeNode) error {
	return l.treeDomain.CreateTreeNode(ctx, node)
}

// DeleteTreeNode 删除树节点
func (l *TreeLogic) DeleteTreeNode(ctx context.Context, id int64, updater string) error {
	return l.treeDomain.DeleteTreeNode(ctx, id, updater)
}

// GetChildrenTreeNode 获取子节点列表
func (l *TreeLogic) GetChildrenTreeNode(ctx context.Context, pid int32, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	return l.treeDomain.GetChildrenTreeNode(ctx, pid, pageSize, pageNum, cmdbType, status)
}

// UpdateTreeNode 更新树节点信息
func (l *TreeLogic) UpdateTreeNode(ctx context.Context, node *model.TreeNode) error {
	return l.treeDomain.UpdateTreeNode(ctx, node)
}

// SyncCMDB 同步CMDB资源数据到树节点
func (l *TreeLogic) SyncCMDB(ctx context.Context, node *model.TreeNode) error {
	return l.treeDomain.SyncCMDB(ctx, node)
}
