package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type AlertManagerPoolLogic struct {
	ctx    context.Context
	domain *domain.AlterManagerPoolDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertManagerPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertManagerPoolLogic {
	return &AlertManagerPoolLogic{
		ctx:    ctx,
		domain: domain.NewAlterManagerPoolDomain(svcCtx),
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (a *AlertManagerPoolLogic) GetMonitorAlertManagerPoolList(ctx context.Context) (*types.GetAlertManagerPoolListResponse, error) {
	pools, err := a.domain.GetMonitorAlertManagerPoolList(ctx, nil)
	if err != nil {
		a.Logger.Errorf("获取 Alertmanager 集群池列表失败: %v", err)
		return nil, err
	}

	return &types.GetAlertManagerPoolListResponse{
		Code:    0,
		Message: "获取 Alertmanager 集群池列表成功",
		Data:    a.domain.BuildAlertManagerPoolRespModel(pools),
	}, nil
}

func (a *AlertManagerPoolLogic) CreateMonitorAlertManagerPool(ctx context.Context, req *types.CreateMonitorAlertManagerPoolRequest) (*types.CreateMonitorAlertManagerPoolResponse, error) {
	// 创建 Alertmanger 集群池
	pool := a.domain.BuildMonitorAlertManagerPoolModel(req.Pool)
	err := a.domain.CreateMonitorAlertManagerPool(ctx, pool)
	if err != nil {
		a.Logger.Errorf("创建 Alertmanager 集群池失败: %v", err)
		return nil, err
	}

	// TODO 更新实例缓存

	a.Logger.Infof("创建 Alertmanager 集群池成功: %+v", pool)

	return &types.CreateMonitorAlertManagerPoolResponse{
		Code:    0,
		Message: "创建 Alertmanager 集群池成功",
	}, nil
}

func (a *AlertManagerPoolLogic) UpdateMonitorAlertManagerPool(ctx context.Context, req *types.UpdateMonitorAlertManagerPoolRequest) (*types.UpdateMonitorAlertManagerPoolResponse, error) {
	// 更新 Alertmanger 集群池
	pool := a.domain.BuildMonitorAlertManagerPoolModel(req.Pool)
	err := a.domain.UpdateMonitorAlertManagerPool(ctx, pool)
	if err != nil {
		a.Logger.Errorf("更新 Alertmanager 集群池失败: %v", err)
		return nil, err
	}

	// TODO 更新实例缓存

	a.Logger.Infof("更新 Alertmanager 集群池成功: %+v", pool)

	return &types.UpdateMonitorAlertManagerPoolResponse{
		Code:    0,
		Message: "更新 Alertmanager 集群池成功",
	}, nil
}

func (a *AlertManagerPoolLogic) DeleteMonitorAlertManagerPool(ctx context.Context, req *types.DeleteMonitorAlertManagerPoolRequest) (*types.DeleteMonitorAlertManagerPoolResponse, error) {
	// TODO 检查 Alertmanager 是否关联发送组

	// 删除 Alertmanger 集群池
	err := a.domain.DeleteMonitorAlertManagerPool(ctx, req.Id)
	if err != nil {
		a.Logger.Errorf("删除 Alertmanager 集群池失败: %v", err)
		return nil, err
	}

	// TODO 更新实例缓存

	a.Logger.Infof("删除 Alertmanager 集群池成功: %+v", req.Id)

	return &types.DeleteMonitorAlertManagerPoolResponse{
		Code:    0,
		Message: "删除 Alertmanager 集群池成功",
	}, nil
}
