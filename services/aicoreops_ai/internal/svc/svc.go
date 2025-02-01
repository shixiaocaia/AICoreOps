package svc

import (
	"sync"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/config"
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
	Mutex     *sync.RWMutex
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)
	store := pkg.InitQdrantStore(c.Qdrant, llm)
	db := pkg.InitDB(c.MySQL)

	return &ServiceContext{
		Config:    c,
		LLM:       llm,
		Qdrant:    store,
		DB:        db,
		MemoryBuf: make(map[string]*memory.ConversationTokenBuffer),
		Mutex:     &sync.RWMutex{},
	}
}
