package svc

import (
	"aicoreops_ai/internal/config"
	"aicoreops_ai/internal/domain"
	"aicoreops_ai/internal/pkg"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	LLM      *ollama.LLM
	Executor *agents.Executor
	Qdrant   *qdrant.Store
	DB       *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)

	store, err := domain.InitQdrantStore(c.Qdrant, llm)
	if err != nil {
		panic(err)
	}

	agent := agents.NewConversationalAgent(llm, []tools.Tool{})
	executor := agents.NewExecutor(agent)

	db, err := pkg.InitGorm(c.MySQL)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		LLM:      llm,
		Executor: executor,
		Qdrant:   store,
		DB:       db,
	}
}
