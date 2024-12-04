package domain

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

// InitQdrantStore 初始化并配置Qdrant向量存储
func InitQdrantStore(documents []schema.Document, embedder *embeddings.EmbedderImpl, serverURL, collectionName string) (*qdrant.Store, error) {
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return nil, fmt.Errorf("解析URL失败: %w", err)
	}

	// 创建Qdrant存储实例
	vectorStore, err := qdrant.New(
		qdrant.WithURL(*parsedURL),
		qdrant.WithCollectionName(collectionName),
		qdrant.WithEmbedder(embedder),
	)
	if err != nil {
		return nil, fmt.Errorf("创建Qdrant存储实例失败: %w", err)
	}

	// 添加文档到存储中
	if len(documents) > 0 {
		if _, err = vectorStore.AddDocuments(context.Background(), documents); err != nil {
			return nil, fmt.Errorf("添加文档失败: %w", err)
		}
	}

	return &vectorStore, nil
}

// RetrieveRelevantDocs 从Qdrant存储中检索相关文档
func RetrieveRelevantDocs(vectorStore *qdrant.Store, queryText string) ([]schema.Document, error) {
	const (
		scoreThreshold = 0.30 // 相似度阈值
		topK          = 5     // 返回最相关文档数量
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
func TextToChunks(dirFile string) ([]schema.Document, error) {
	const (
		chunkSize    = 768
		chunkOverlap = 64
	)

	file, err := os.Open(dirFile)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	docLoaded := documentloaders.NewText(file)
	split := textsplitter.NewRecursiveCharacter()
	split.ChunkSize = chunkSize
	split.ChunkOverlap = chunkOverlap

	docs, err := docLoaded.LoadAndSplit(context.Background(), split)
	if err != nil {
		return nil, fmt.Errorf("拆分文本失败: %w", err)
	}

	return docs, nil
}
