package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type AlertmanagerLogic struct {
	ctx    context.Context
	domain *domain.AltermanagerDomain
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertmanagerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertmanagerLogic {
	return &AlertmanagerLogic{
		ctx:    ctx,
		domain: domain.NewAltermanagerDomain(svcCtx),
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (a *AlertmanagerLogic) GetMonitorAlertmanagerPoolList(ctx context.Context) (*types.GetAlertmanagerPoolListResponse, error) {
	pools, err := a.domain.GetMonitorAlertmanagerPoolList(ctx, nil)
	if err != nil {
		a.Logger.Errorf("获取 Alertmanager 集群池列表失败: %v", err)
		return nil, err
	}

	return &types.GetAlertmanagerPoolListResponse{
		Code:    0,
		Message: "获取 Alertmanager 集群池列表成功",
		Data:    a.domain.BuildAlertmanagerPoolRespModel(pools),
	}, nil
}

func (a *AlertmanagerLogic) CreateMonitorAlertManagerPool(ctx context.Context, req *types.CreateMonitorAlertManagerPoolRequest) (*types.CreateMonitorAlertManagerPoolResponse, error) {
	// 创建 Alertmanger 集群池
	pool := a.domain.BuildMonitorAlertmanagerPoolModel(req.Pool)
	err := a.domain.CreateMonitorAlertmanagerPool(ctx, pool)
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

func (a *AlertmanagerLogic) UpdateMonitorAlertManagerPool(ctx context.Context, req *types.UpdateMonitorAlertManagerPoolRequest) (*types.UpdateMonitorAlertManagerPoolResponse, error) {
	// 更新 Alertmanger 集群池
	pool := a.domain.BuildMonitorAlertmanagerPoolModel(req.Pool)
	err := a.domain.UpdateMonitorAlertmanagerPool(ctx, pool)
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

func (a *AlertmanagerLogic) DeleteMonitorAlertManagerPool(ctx context.Context, req *types.DeleteMonitorAlertManagerPoolRequest) (*types.DeleteMonitorAlertManagerPoolResponse, error) {
	// TODO 检查 Alertmanager 是否关联发送组

	// 删除 Alertmanger 集群池
	err := a.domain.DeleteMonitorAlertmanagerPool(ctx, req.Id)
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
