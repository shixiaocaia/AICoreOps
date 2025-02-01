package pkg

import (
	"fmt"
	"net/url"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/config"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

// InitQdrantStore 初始化并配置Qdrant向量存储
func InitQdrantStore(c config.QdrantConfig, llm *ollama.LLM) *qdrant.Store {
	parsedURL, err := url.Parse(c.Url)
	if err != nil {
		panic(fmt.Errorf("解析Qdrant URL失败: %w", err))
	}

	ollamaEmbedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		panic(fmt.Errorf("创建Ollama嵌入器失败: %w", err))
	}

	vectorStore, err := qdrant.New(
		qdrant.WithURL(*parsedURL),
		qdrant.WithCollectionName(c.CollectionName),
		qdrant.WithEmbedder(ollamaEmbedder),
	)
	if err != nil {
		panic(fmt.Errorf("创建Qdrant向量存储失败: %w", err))
	}

	return &vectorStore
}
