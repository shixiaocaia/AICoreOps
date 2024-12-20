package domain

import (
	"aicoreops_tree/internal/dao"
	"aicoreops_tree/internal/model"
	"aicoreops_tree/internal/repo"
	"context"

	"gorm.io/gorm"
)

type TreeDomain struct {
	repo repo.TreeRepo
}

func NewTreeDomain(db *gorm.DB) *TreeDomain {
	return &TreeDomain{
		repo: dao.NewTreeDao(db),
	}
}

// ListTreeNode 查询树节点列表
func (d *TreeDomain) ListTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status, keyword string) ([]*model.TreeNode, int32, error) {
	return d.repo.ListTreeNode(ctx, pageSize, pageNum, cmdbType, status, keyword)
}

// SelectTreeNode 查询单个树节点
func (d *TreeDomain) SelectTreeNode(ctx context.Context, id int64, cmdbID string) (*model.TreeNode, error) {
	return d.repo.SelectTreeNode(ctx, id, cmdbID)
}

// GetTopTreeNode 查询顶级树节点
func (d *TreeDomain) GetTopTreeNode(ctx context.Context, limit int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	return d.repo.GetTopTreeNode(ctx, limit, cmdbType, status)
}

// ListLeafTreeNode 查询叶子节点列表
func (d *TreeDomain) ListLeafTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	return d.repo.ListLeafTreeNode(ctx, pageSize, pageNum, cmdbType, status)
}

// CreateTreeNode 创建树节点
func (d *TreeDomain) CreateTreeNode(ctx context.Context, node *model.TreeNode) error {
	return d.repo.CreateTreeNode(ctx, node)
}

// DeleteTreeNode 删除树节点
func (d *TreeDomain) DeleteTreeNode(ctx context.Context, id int64, updater string) error {
	return d.repo.DeleteTreeNode(ctx, id, updater)
}

// GetChildrenTreeNode 获取子节点列表
func (d *TreeDomain) GetChildrenTreeNode(ctx context.Context, pid int32, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	return d.repo.GetChildrenTreeNode(ctx, pid, pageSize, pageNum, cmdbType, status)
}

// UpdateTreeNode 更新树节点信息
func (d *TreeDomain) UpdateTreeNode(ctx context.Context, node *model.TreeNode) error {
	return d.repo.UpdateTreeNode(ctx, node)
}

// SyncCMDB 同步CMDB资源数据到树节点
func (d *TreeDomain) SyncCMDB(ctx context.Context, node *model.TreeNode) error {
	return d.repo.SyncCMDB(ctx, node)
}
