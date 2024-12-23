package domain

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/config"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

// InitQdrantStore 初始化并配置Qdrant向量存储
func InitQdrantStore(c config.QdrantConfig, llm *ollama.LLM) (*qdrant.Store, error) {
	parsedURL, err := url.Parse(c.Url)
	if err != nil {
		return nil, fmt.Errorf("解析URL失败: %w", err)
	}

	ollamaEmbedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return nil, fmt.Errorf("创建文档嵌入器失败: %w", err)
	}

	vectorStore, err := qdrant.New(
		qdrant.WithURL(*parsedURL),
		qdrant.WithCollectionName(c.CollectionName),
		qdrant.WithEmbedder(ollamaEmbedder),
	)
	if err != nil {
		return nil, fmt.Errorf("创建Qdrant存储实例失败: %w", err)
	}

	return &vectorStore, nil
}

func InsertDocsToQdrantStore(vectorStore *qdrant.Store, docs []schema.Document) error {
	_, err := vectorStore.AddDocuments(context.Background(), docs)
	if err != nil {
		return fmt.Errorf("添加文档失败: %w", err)
	}

	return nil
}

// RetrieveRelevantDocs 从Qdrant存储中检索相关文档
func RetrieveRelevantDocs(vectorStore *qdrant.Store, queryText string) ([]schema.Document, error) {
	const (
		scoreThreshold = 0.30 // 相似度阈值
		topK           = 5    // 返回最相关文档数量
	)

	// 设置向量检索选项
	retrievalOptions := []vectorstores.Option{
		vectorstores.WithScoreThreshold(scoreThreshold),
	}

	// 创建检索器实例并执行检索
	docRetriever := vectorstores.ToRetriever(vectorStore, topK, retrievalOptions...)
	relevantDocs, err := docRetriever.GetRelevantDocuments(context.Background(), queryText)
	if err != nil {
		return nil, fmt.Errorf("检索文档失败: %w", err)
	}

	return relevantDocs, nil
}

// TextToChunks 将文本文件拆分成多个文档块
func TextToChunks(contentBytes []byte) ([]schema.Document, error) {
	const (
		chunkSize    = 768
		chunkOverlap = 64
	)

	contentReader := bytes.NewReader(contentBytes)
	docLoaded := documentloaders.NewText(contentReader)

	splitter := textsplitter.NewRecursiveCharacter()
	splitter.ChunkSize = chunkSize
	splitter.ChunkOverlap = chunkOverlap

	docs, err := docLoaded.LoadAndSplit(context.Background(), splitter)
	if err != nil {
		return nil, fmt.Errorf("拆分文本失败: %w", err)
	}

	return docs, nil
}
