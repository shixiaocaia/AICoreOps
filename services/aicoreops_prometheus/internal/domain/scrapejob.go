package domain

import (
	"context"
	"encoding/json"
	"errors"

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
		repo: dao.NewScrapeJobDao(svcCtx.DB),
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
	// 将 tree_node_ids 转换为 JSON 字符串
	treeNodeIDsJSON, _ := json.Marshal(job.TreeNodeIds)

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
		TreeNodeIDs:              string(treeNodeIDsJSON),
		KubeConfigFilePath:       job.KubeConfigFilePath,
		TlsCaFilePath:            job.TlsCaFilePath,
		TlsCaContent:             job.TlsCaContent,
		BearerToken:              job.BearerToken,
		BearerTokenFile:          job.BearerTokenFile,
		KubernetesSdRole:         job.KubernetesSdRole,
	}
}

func (d *ScrapeJobDomain) BuildMonitorScrapeJobRespModel(jobs []*model.MonitorScrapeJob) []*types.ScrapeJob {
	var result []*types.ScrapeJob
	for _, job := range jobs {
		// 将 JSON 字符串转换为 []int64
		var treeNodeIDs []int64
		_ = json.Unmarshal([]byte(job.TreeNodeIDs), &treeNodeIDs)

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
			TreeNodeIds:              treeNodeIDs,
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
