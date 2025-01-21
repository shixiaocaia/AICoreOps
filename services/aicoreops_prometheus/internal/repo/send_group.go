package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
)

// SendGroupRepo 发送组Repo
type SendGroupRepo interface {
	GetMonitorSendGroupByPoolId(ctx context.Context, poolId int64) ([]*model.MonitorSendGroup, error)
	GetMonitorSendGroupByOnDutyGroupId(ctx context.Context, onDutyGroupID int64) ([]*model.MonitorSendGroup, error)
	SearchMonitorSendGroupByName(ctx context.Context, name string) ([]*model.MonitorSendGroup, error)
	GetMonitorSendGroupList(ctx context.Context) ([]*model.MonitorSendGroup, error)
	GetMonitorSendGroupById(ctx context.Context, id int64) (*model.MonitorSendGroup, error)
	CreateMonitorSendGroup(ctx context.Context, monitorSendGroup *model.MonitorSendGroup) error
	UpdateMonitorSendGroup(ctx context.Context, monitorSendGroup *model.MonitorSendGroup) error
	DeleteMonitorSendGroup(ctx context.Context, id int64) error
	CheckMonitorSendGroupExists(ctx context.Context, sendGroup *model.MonitorSendGroup) (bool, error)
	CheckMonitorSendGroupNameExists(ctx context.Context, sendGroup *model.MonitorSendGroup) (bool, error)
}
