package dao

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"gorm.io/gorm"
)

type SendGroupDao struct {
	db *gorm.DB
}

func NewSendGroupDao(db *gorm.DB) *SendGroupDao {
	return &SendGroupDao{db: db}
}

// GetMonitorSendGroupByPoolId 获取指定AlertManager池下的发送组
func (d *SendGroupDao) GetMonitorSendGroupByPoolId(ctx context.Context, poolId int64) ([]*model.MonitorSendGroup, error) {
	var sendGroups []*model.MonitorSendGroup
	if err := d.db.WithContext(ctx).Where("pool_id = ?", poolId).Find(&sendGroups).Error; err != nil {
		return nil, err
	}
	return sendGroups, nil
}

// GetMonitorSendGroupByOnDutyGroupId 获取指定值班组下的发送组
func (d *SendGroupDao) GetMonitorSendGroupByOnDutyGroupId(ctx context.Context, onDutyGroupID int64) ([]*model.MonitorSendGroup, error) {
	var sendGroups []*model.MonitorSendGroup
	if err := d.db.WithContext(ctx).Where("on_duty_group_id = ?", onDutyGroupID).Find(&sendGroups).Error; err != nil {
		return nil, err
	}
	return sendGroups, nil
}

// SearchMonitorSendGroupByName 根据名称模糊搜索发送组
func (d *SendGroupDao) SearchMonitorSendGroupByName(ctx context.Context, name string) ([]*model.MonitorSendGroup, error) {
	var sendGroups []*model.MonitorSendGroup
	if err := d.db.WithContext(ctx).Where("name_zh LIKE ?", "%"+name+"%").Find(&sendGroups).Error; err != nil {
		return nil, err
	}
	return sendGroups, nil
}

// GetMonitorSendGroupList 获取所有发送组
func (d *SendGroupDao) GetMonitorSendGroupList(ctx context.Context) ([]*model.MonitorSendGroup, error) {
	var sendGroups []*model.MonitorSendGroup
	if err := d.db.WithContext(ctx).Find(&sendGroups).Error; err != nil {
		return nil, err
	}
	return sendGroups, nil
}

// GetMonitorSendGroupById 获取指定ID的发送组
func (d *SendGroupDao) GetMonitorSendGroupById(ctx context.Context, id int64) (*model.MonitorSendGroup, error) {
	var sendGroup model.MonitorSendGroup
	if err := d.db.WithContext(ctx).Where("id = ?", id).First(&sendGroup).Error; err != nil {
		return nil, err
	}
	return &sendGroup, nil
}

// CreateMonitorSendGroup 创建发送组
func (d *SendGroupDao) CreateMonitorSendGroup(ctx context.Context, monitorSendGroup *model.MonitorSendGroup) error {
	return d.db.WithContext(ctx).Create(monitorSendGroup).Error
}

// UpdateMonitorSendGroup 更新发送组
func (d *SendGroupDao) UpdateMonitorSendGroup(ctx context.Context, monitorSendGroup *model.MonitorSendGroup) error {
	return d.db.WithContext(ctx).Save(monitorSendGroup).Error
}

// DeleteMonitorSendGroup 删除发送组
func (d *SendGroupDao) DeleteMonitorSendGroup(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Delete(&model.MonitorSendGroup{}, id).Error
}

// CheckMonitorSendGroupExists 检查发送组是否存在
func (d *SendGroupDao) CheckMonitorSendGroupExists(ctx context.Context, sendGroup *model.MonitorSendGroup) (bool, error) {
	var count int64
	if err := d.db.WithContext(ctx).Model(&model.MonitorSendGroup{}).Where("name_zh = ?", sendGroup.NameZh).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckMonitorSendGroupNameExists 检查发送组名称是否存在
func (d *SendGroupDao) CheckMonitorSendGroupNameExists(ctx context.Context, sendGroup *model.MonitorSendGroup) (bool, error) {
	var count int64
	if err := d.db.WithContext(ctx).Model(&model.MonitorSendGroup{}).Where("name_zh = ?", sendGroup.NameZh).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetMonitorSendGroupByName 根据名称获取发送组
func (d *SendGroupDao) GetMonitorSendGroupByName(ctx context.Context, name string) (*model.MonitorSendGroup, error) {
	var sendGroup model.MonitorSendGroup
	if err := d.db.WithContext(ctx).Where("name = ?", name).First(&sendGroup).Error; err != nil {
		return nil, err
	}
	return &sendGroup, nil
}
