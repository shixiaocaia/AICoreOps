package svc

import (
	"aicoreops_ai/internal/config"
	"aicoreops_ai/internal/pkg"

	"github.com/tmc/langchaingo/llms/ollama"
)

type ServiceContext struct {
	Config   config.Config
	LLM      *ollama.LLM
	Embedder *ollama.LLM
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)
	embedder := pkg.InitEmbedderLLM(c.Qdrant)
	return &ServiceContext{
		Config:   c,
		LLM:      llm,
		Embedder: embedder,
	}
}
