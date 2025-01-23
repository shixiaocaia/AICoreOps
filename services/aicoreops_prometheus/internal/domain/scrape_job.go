package domain

import (
	"context"
	"errors"
	"strconv"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/repo"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_prometheus/types"
)

type ScrapeJobDomain struct {
	repo repo.MonitorScrapeJobRepo
}

func NewScrapeJobDomain(svcCtx *svc.ServiceContext) *ScrapeJobDomain {
	return &ScrapeJobDomain{
		repo: dao.NewMonitorScrapeJobDAO(svcCtx.DB),
	}
}

func (d *ScrapeJobDomain) GetMonitorScrapeJobList(ctx context.Context) ([]*model.MonitorScrapeJob, error) {
	return d.repo.GetMonitorScrapeJobList(ctx)
}

func (d *ScrapeJobDomain) CreateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error {
	// 检查 job 是否存在
	jobs, err := d.repo.SearchMonitorScrapeJobByName(ctx, job.Name)
	if err != nil {
		return err
	}
	if len(jobs) > 0 {
		return errors.New("job already exists")
	}

	return d.repo.CreateMonitorScrapeJob(ctx, job)
}

func (d *ScrapeJobDomain) UpdateMonitorScrapeJob(ctx context.Context, job *model.MonitorScrapeJob) error {
	return d.repo.UpdateMonitorScrapeJob(ctx, job)
}

func (d *ScrapeJobDomain) DeleteMonitorScrapeJob(ctx context.Context, id int64) error {
	return d.repo.DeleteMonitorScrapeJob(ctx, id)
}

func (d *ScrapeJobDomain) BuildMonitorScrapeJobModel(job *types.ScrapeJob) *model.MonitorScrapeJob {
	ids := make([]string, len(job.TreeNodeIds))
	for i, id := range job.TreeNodeIds {
		ids[i] = strconv.FormatInt(id, 10)
	}
	return &model.MonitorScrapeJob{
		ID:                       job.Id,
		Name:                     job.Name,
		UserID:                   job.UserId,
		Enable:                   job.Enable,
		ServiceDiscoveryType:     job.ServiceDiscoveryType,
		MetricsPath:              job.MetricsPath,
		Scheme:                   job.Scheme,
		ScrapeInterval:           job.ScrapeInterval,
		ScrapeTimeout:            job.ScrapeTimeout,
		PoolID:                   job.PoolId,
		RelabelConfigsYamlString: job.RelabelConfigsYamlString,
		RefreshInterval:          job.RefreshInterval,
		Port:                     job.Port,
		TreeNodeIDs:              model.StringList(ids),
		KubeConfigFilePath:       job.KubeConfigFilePath,
		TlsCaFilePath:            job.TlsCaFilePath,
		TlsCaContent:             job.TlsCaContent,
		BearerToken:              job.BearerToken,
		BearerTokenFile:          job.BearerTokenFile,
		KubernetesSdRole:         job.KubernetesSdRole,
	}
}

func (d *ScrapeJobDomain) BuildScrapeJobRespModel(jobs []*model.MonitorScrapeJob) []*types.ScrapeJob {
	var result []*types.ScrapeJob
	for _, job := range jobs {
		ids := make([]int64, len(job.TreeNodeIDs))
		for i, id := range job.TreeNodeIDs {
			ids[i], _ = strconv.ParseInt(id, 10, 64)
		}
		result = append(result, &types.ScrapeJob{
			Id:                       job.ID,
			Name:                     job.Name,
			UserId:                   job.UserID,
			Enable:                   job.Enable,
			ServiceDiscoveryType:     job.ServiceDiscoveryType,
			MetricsPath:              job.MetricsPath,
			Scheme:                   job.Scheme,
			ScrapeInterval:           job.ScrapeInterval,
			ScrapeTimeout:            job.ScrapeTimeout,
			PoolId:                   job.PoolID,
			RelabelConfigsYamlString: job.RelabelConfigsYamlString,
			RefreshInterval:          job.RefreshInterval,
			Port:                     job.Port,
			TreeNodeIds:              ids,
			KubeConfigFilePath:       job.KubeConfigFilePath,
			TlsCaFilePath:            job.TlsCaFilePath,
			TlsCaContent:             job.TlsCaContent,
			BearerToken:              job.BearerToken,
			BearerTokenFile:          job.BearerTokenFile,
			KubernetesSdRole:         job.KubernetesSdRole,
		})
	}
	return result
}
