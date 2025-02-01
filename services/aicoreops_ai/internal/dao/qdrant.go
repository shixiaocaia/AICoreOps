package dao

import (
	"bytes"
	"context"
	"fmt"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

// QdrantRepository 定义向量存储库接口
type QdrantRepository interface {
	// TextToDocuments 将文本转换为文档
	TextToDocuments(ctx context.Context, content []byte) ([]schema.Document, error)

	// StoreDocuments 存储文档到向量库
	StoreDocuments(ctx context.Context, docs []schema.Document) error

	// SearchSimilarDocuments 搜索相似文档
	SearchSimilarDocuments(ctx context.Context, query string, opts ...SearchOption) ([]schema.Document, error)
}

// SearchOption 定义搜索选项
type SearchOption func(*searchOptions)

type searchOptions struct {
	scoreThreshold float32
	topK           int
}

// WithScoreThreshold 设置相似度阈值选项
func WithScoreThreshold(threshold float32) SearchOption {
	return func(o *searchOptions) {
		o.scoreThreshold = threshold
	}
}

// WithTopK 设置返回结果数量选项
func WithTopK(k int) SearchOption {
	return func(o *searchOptions) {
		o.topK = k
	}
}

// QdrantDAO 实现向量存储库接口
type QdrantDAO struct {
	store *qdrant.Store
	opts  *options
}

type options struct {
	chunkSize    int
	chunkOverlap int
}

// NewQdrantDAO 创建 QdrantDAO 实例
func NewQdrantDAO(store *qdrant.Store, opts ...Option) *QdrantDAO {
	defaultOpts := &options{
		chunkSize:    768,
		chunkOverlap: 64,
	}

	for _, opt := range opts {
		opt(defaultOpts)
	}

	return &QdrantDAO{
		store: store,
		opts:  defaultOpts,
	}
}

// Option 定义 DAO 配置选项
type Option func(*options)

// WithChunkSize 设置文档分块大小
func WithChunkSize(size int) Option {
	return func(o *options) {
		o.chunkSize = size
	}
}

// WithChunkOverlap 设置分块重叠大小
func WithChunkOverlap(overlap int) Option {
	return func(o *options) {
		o.chunkOverlap = overlap
	}
}

// TextToDocuments 实现文本转文档接口
func (q *QdrantDAO) TextToDocuments(ctx context.Context, content []byte) ([]schema.Document, error) {
	contentReader := bytes.NewReader(content)
	docLoader := documentloaders.NewText(contentReader)

	splitter := textsplitter.NewRecursiveCharacter()
	splitter.ChunkSize = q.opts.chunkSize
	splitter.ChunkOverlap = q.opts.chunkOverlap

	docs, err := docLoader.LoadAndSplit(ctx, splitter)
	if err != nil {
		return nil, fmt.Errorf("split text failed: %w", err)
	}

	return docs, nil
}

// DocumentMetadata 定义文档元数据
type DocumentMetadata struct {
	Title string `json:"title"`
}

// StoreDocumentWithMetadata 存储带元数据的文档
func (q *QdrantDAO) StoreDocumentWithMetadata(ctx context.Context, docs []schema.Document, title string) error {
	for i := range docs {
		if docs[i].Metadata == nil {
			docs[i].Metadata = make(map[string]any)
		}
		docs[i].Metadata["title"] = title
	}

	if _, err := q.store.AddDocuments(ctx, docs); err != nil {
		return fmt.Errorf("store documents failed: %w", err)
	}

	return nil
}

// SearchSimilarDocuments 实现搜索相似文档接口
func (q *QdrantDAO) SearchSimilarDocuments(ctx context.Context, title, query string, opts ...SearchOption) ([]schema.Document, error) {
	options := &searchOptions{
		scoreThreshold: 0.3,
		topK:           5,
	}

	for _, opt := range opts {
		opt(options)
	}

	// 添加元数据过滤选项
	retrievalOpts := []vectorstores.Option{
		vectorstores.WithScoreThreshold(options.scoreThreshold),
	}
	if title != "" {
		retrievalOpts = append(retrievalOpts, vectorstores.WithFilters(map[string]interface{}{
			"title": title,
		}))
	}

	retriever := vectorstores.ToRetriever(q.store, options.topK, retrievalOpts...)
	docs, err := retriever.GetRelevantDocuments(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("search documents failed: %w", err)
	}

	return docs, nil
}
