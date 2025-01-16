package svc

import (
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/pkg"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	LLM                 *ollama.LLM
	Qdrant              *qdrant.Store
	HistoryModel        model.HistoryModel
	HistorySessionModel model.HistorySessionModel
	MemoryBuf           map[string]*memory.ConversationTokenBuffer
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)

	store, err := domain.InitQdrantStore(c.Qdrant, llm)
	if err != nil {
		panic(err)
	}

	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:              c,
		LLM:                 llm,
		Qdrant:              store,
		HistoryModel:        model.NewHistoryModel(conn),
		HistorySessionModel: model.NewHistorySessionModel(conn),
		MemoryBuf:           make(map[string]*memory.ConversationTokenBuffer),
	}
}
