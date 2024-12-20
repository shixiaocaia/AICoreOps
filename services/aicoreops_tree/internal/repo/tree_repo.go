package repo

import (
	"aicoreops_tree/internal/model"
	"context"
)

// TreeRepo 定义资源树仓储层接口
type TreeRepo interface {
	ListTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status, keyword string) ([]*model.TreeNode, int32, error)
	SelectTreeNode(ctx context.Context, id int64, cmdbID string) (*model.TreeNode, error)
	GetTopTreeNode(ctx context.Context, limit int32, cmdbType, status string) ([]*model.TreeNode, int32, error)
	ListLeafTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error)
	CreateTreeNode(ctx context.Context, node *model.TreeNode) error
	DeleteTreeNode(ctx context.Context, id int64, updater string) error
	GetChildrenTreeNode(ctx context.Context, pid int32, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error)
	UpdateTreeNode(ctx context.Context, node *model.TreeNode) error
	SyncCMDB(ctx context.Context, node *model.TreeNode) error
}
