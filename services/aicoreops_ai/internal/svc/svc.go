package svc

import (
	"aicoreops_ai/internal/config"
	"aicoreops_ai/internal/domain"
	"aicoreops_ai/internal/pkg"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

type ServiceContext struct {
	Config config.Config
	LLM    *ollama.LLM
	Qdrant *qdrant.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)
	store, err := domain.InitQdrantStore(c.Qdrant)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		LLM:    llm,
		Qdrant: store,
	}
}
