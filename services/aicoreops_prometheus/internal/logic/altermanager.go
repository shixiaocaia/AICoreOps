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

func (a *AlertmanagerLogic) GetMonitorAlertmanagerPoolList(ctx context.Context, req *types.GetAlertmanagerPoolListRequest) (*types.GetAlertmanagerPoolListResponse, error) {
	panic("not implemented")
}

func (a *AlertmanagerLogic) CreateMonitorAlertManagerPool(ctx context.Context, req *types.CreateMonitorAlertManagerPoolRequest) (*types.CreateMonitorAlertManagerPoolResponse, error) {
	// 创建 Alertmanger 集群池
	pool := a.domain.BuildMonitorAlertmanagerPoolModel(ctx, req.Pool)
	err := a.domain.CreateMonitorAlertmanagerPool(ctx, pool)
	if err != nil {
		return nil, err
	}

	// 更新实例缓存

	return &types.CreateMonitorAlertManagerPoolResponse{
		Code:    0,
		Message: "创建 Alertmanager 集群池成功",
	}, nil
}

func (a *AlertmanagerLogic) UpdateMonitorAlertManagerPool(ctx context.Context, req *types.UpdateMonitorAlertManagerPoolRequest) (*types.UpdateMonitorAlertManagerPoolResponse, error) {
	panic("not implemented")
}

func (a *AlertmanagerLogic) DeleteMonitorAlertManagerPool(ctx context.Context, req *types.DeleteMonitorAlertManagerPoolRequest) (*types.DeleteMonitorAlertManagerPoolResponse, error) {
	panic("not implemented")
}
