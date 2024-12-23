package dao

import (
	"aicoreops_tree/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

type TreeDao struct {
	db *gorm.DB
}

func NewTreeDao(db *gorm.DB) *TreeDao {
	return &TreeDao{
		db: db,
	}
}

// ListTreeNode 查询树节点列表
func (t *TreeDao) ListTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status, keyword string) ([]*model.TreeNode, int32, error) {
	if pageSize <= 0 || pageNum <= 0 {
		return nil, 0, ErrInvalidRequest
	}

	var nodes []*model.TreeNode
	var total int64

	// 构建基础查询
	query := t.db.WithContext(ctx).Model(&model.TreeNode{}).Where("is_deleted = ?", 0)

	// 添加过滤条件
	if cmdbType != "" {
		query = query.Where("cmdb_type = ?", cmdbType)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if keyword != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (pageNum - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&nodes).Error; err != nil {
		return nil, 0, err
	}

	return nodes, int32(total), nil
}

// SelectTreeNode 查询单个树节点
func (t *TreeDao) SelectTreeNode(ctx context.Context, id int64, cmdbID string) (*model.TreeNode, error) {
	if id <= 0 && cmdbID == "" {
		return nil, ErrInvalidRequest
	}

	var node model.TreeNode

	query := t.db.WithContext(ctx).Where("is_deleted = ?", 0)

	// 优先使用ID查询
	if id > 0 {
		query = query.Where("id = ?", id)
	} else {
		query = query.Where("cmdb_id = ?", cmdbID)
	}

	if err := query.First(&node).Error; err != nil {
		return nil, err
	}

	return &node, nil
}

// GetTopTreeNode 查询顶级树节点
func (t *TreeDao) GetTopTreeNode(ctx context.Context, limit int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	var nodes []*model.TreeNode
	var total int64

	// 查询顶级节点(pid=0)
	query := t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("pid = ? AND is_deleted = ?", 0, 0)

	if cmdbType != "" {
		query = query.Where("cmdb_type = ?", cmdbType)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(int(limit))
	}

	if err := query.Find(&nodes).Error; err != nil {
		return nil, 0, err
	}

	return nodes, int32(total), nil
}

// ListLeafTreeNode 查询叶子节点列表
func (t *TreeDao) ListLeafTreeNode(ctx context.Context, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	if pageSize <= 0 || pageNum <= 0 {
		return nil, 0, ErrInvalidRequest
	}

	var nodes []*model.TreeNode
	var total int64

	// 查询叶子节点(is_leaf=1)
	query := t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("is_leaf = ? AND is_deleted = ?", 1, 0)

	if cmdbType != "" {
		query = query.Where("cmdb_type = ?", cmdbType)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&nodes).Error; err != nil {
		return nil, 0, err
	}

	return nodes, int32(total), nil
}

// CreateTreeNode 创建树节点
func (t *TreeDao) CreateTreeNode(ctx context.Context, node *model.TreeNode) error {
	if node == nil || node.Title == "" {
		return ErrInvalidRequest
	}

	// 设置创建和更新时间
	node.CreateTime = time.Now().Unix()
	node.UpdateTime = time.Now().Unix()

	return t.db.WithContext(ctx).Create(node).Error
}

// DeleteTreeNode 删除树节点
func (t *TreeDao) DeleteTreeNode(ctx context.Context, id int64, updater string) error {
	if id <= 0 {
		return ErrInvalidRequest
	}

	// 检查是否存在子节点
	var count int64
	if err := t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("pid = ? AND is_deleted = ?", id, 0).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("节点存在子节点,不允许删除")
	}

	// 执行软删除
	updates := map[string]interface{}{
		"is_deleted":  1,
		"updater":     updater,
		"update_time": time.Now().Unix(),
	}
	return t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// GetChildrenTreeNode 获取子节点列表
func (t *TreeDao) GetChildrenTreeNode(ctx context.Context, pid int32, pageSize, pageNum int32, cmdbType, status string) ([]*model.TreeNode, int32, error) {
	if pid < 0 || pageSize <= 0 || pageNum <= 0 {
		return nil, 0, ErrInvalidRequest
	}

	var nodes []*model.TreeNode
	var total int64

	// 查询指定父节点的子节点
	query := t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("pid = ? AND is_deleted = ?", pid, 0)

	if cmdbType != "" {
		query = query.Where("cmdb_type = ?", cmdbType)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&nodes).Error; err != nil {
		return nil, 0, err
	}

	return nodes, int32(total), nil
}

// UpdateTreeNode 更新树节点信息
func (t *TreeDao) UpdateTreeNode(ctx context.Context, node *model.TreeNode) error {
	if node == nil || node.ID <= 0 {
		return ErrInvalidRequest
	}

	// 更新节点属性
	updates := map[string]interface{}{
		"title":       node.Title,
		"description": node.Description,
		"is_leaf":     node.IsLeaf,
		"cmdb_attrs":  node.CMDBAttrs,
		"updater":     node.Updater,
		"status":      node.Status,
		"update_time": time.Now().Unix(),
	}

	return t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("id = ? AND is_deleted = ?", node.ID, 0).
		Updates(updates).Error
}

// SyncCMDB 同步CMDB资源数据到树节点
func (t *TreeDao) SyncCMDB(ctx context.Context, node *model.TreeNode) error {
	if node == nil || node.ID <= 0 || node.CMDBID == "" {
		return ErrInvalidRequest
	}

	// TODO: 同步CMDB资源数据到树节点逻辑

	// 更新CMDB相关属性
	updates := map[string]interface{}{
		"cmdb_id":    node.CMDBID,
		"cmdb_type":  node.CMDBType,
		"cmdb_attrs": node.CMDBAttrs,
	}

	return t.db.WithContext(ctx).Model(&model.TreeNode{}).
		Where("id = ? AND is_deleted = ?", node.ID, 0).
		Updates(updates).Error
}
