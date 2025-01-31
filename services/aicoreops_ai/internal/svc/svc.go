package svc

import (
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/config"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/pkg"
	"gorm.io/gorm"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

type ServiceContext struct {
	Config    config.Config
	LLM       *ollama.LLM
	Qdrant    *qdrant.Store
	DB        *gorm.DB
	MemoryBuf map[string]*memory.ConversationTokenBuffer
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)

	store, err := domain.InitQdrantStore(c.Qdrant, llm)
	if err != nil {
		panic(err)
	}

	db := pkg.InitDB(c.MySQL)

	return &ServiceContext{
		Config:    c,
		LLM:       llm,
		Qdrant:    store,
		DB:        db,
		MemoryBuf: make(map[string]*memory.ConversationTokenBuffer),
	}
}
