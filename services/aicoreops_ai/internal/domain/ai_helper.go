package domain

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/dao"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/model"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/repo"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/types"
	"gorm.io/gorm"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

type AIHelperDomain struct {
	HistoryRepo        repo.HistoryRepo
	HistorySessionRepo repo.HistorySessionRepo
	Qdrant             *dao.QdrantDAO
}

func NewAIHelperDomain(db *gorm.DB, qd *qdrant.Store) *AIHelperDomain {
	return &AIHelperDomain{
		HistoryRepo:        dao.NewHistoryDAO(db),
		HistorySessionRepo: dao.NewHistorySessionDAO(db),
		Qdrant:             dao.NewQdrantDAO(qd),
	}
}

type ChatSession struct {
	UserID    int64
	SessionID string
	MemoryBuf *memory.ConversationTokenBuffer
	IsNew     bool
}

func (d *AIHelperDomain) GetHistorySessionList(ctx context.Context, uid int64, limit, offset int) ([]*types.HistorySession, error) {
	list, err := d.HistorySessionRepo.GetHistorySessionList(ctx, uid, offset, limit)
	if err != nil {
		return nil, err
	}

	sessions := make([]*types.HistorySession, 0, len(list))
	for _, h := range list {
		sessions = append(sessions, &types.HistorySession{
			Title:     h.Title,
			SessionId: h.SessionID,
		})
	}

	return sessions, nil
}

func (d *AIHelperDomain) GetHistoryBySessionID(ctx context.Context, sessionID string) ([]*types.GetChatHistoryResponse_ChatMessage, error) {
	histories, err := d.HistoryRepo.GetHistoryBySessionID(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("获取历史记录失败: %w", err)
	}

	return d.BuildHistoryRespModel(ctx, histories), nil
}

func (d *AIHelperDomain) UploadDocument(ctx context.Context, title, content string) error {
	// 1. 解码 Base64 内容
	fileBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return fmt.Errorf("解码 Base64 内容失败: %w", err)
	}

	// 2. 将文本文件拆分成文档块
	docs, err := d.Qdrant.TextToDocuments(ctx, fileBytes)
	if err != nil {
		return fmt.Errorf("拆分文本失败: %w", err)
	}

	// 3. 存储到 Qdrant 向量存储中
	if err = d.Qdrant.StoreDocumentWithMetadata(ctx, docs, title); err != nil {
		return fmt.Errorf("添加文档失败: %w", err)
	}

	return nil
}

func (d *AIHelperDomain) LoadHistoryToMemory(ctx context.Context, sc *svc.ServiceContext, sessionID string, histories []*types.GetChatHistoryResponse_ChatMessage) error {
	if _, ok := sc.MemoryBuf[sessionID]; ok {
		return nil
	}
	buf := memory.NewConversationTokenBuffer(
		sc.LLM,
		10000,
	)
	for _, h := range histories {
		err := buf.SaveContext(ctx, map[string]any{"question": h.Question}, map[string]any{"answer": h.Answer})
		if err != nil {
			return err
		}
	}

	// 并发写入
	sc.Mutex.Lock()
	sc.MemoryBuf[sessionID] = buf
	sc.Mutex.Unlock()

	return nil
}

func (d *AIHelperDomain) CheckSession(ctx context.Context) (userID int64, sessionID string, err error) {
	userID = ctx.Value("userId").(int64)
	if userID == 0 {
		return 0, "", fmt.Errorf("用户ID为空")
	}

	sessionID = ctx.Value("sessionID").(string)
	if sessionID == "" {
		return 0, "", fmt.Errorf("会话ID为空")
	}

	return userID, sessionID, nil
}

func (d *AIHelperDomain) GetMemoryBuf(ctx context.Context, sessionID string, llm *ollama.LLM, mp map[string]*memory.ConversationTokenBuffer, mutex *sync.RWMutex) (*memory.ConversationTokenBuffer, bool, error) {
	buf, ok := mp[sessionID]
	if !ok {
		buf = memory.NewConversationTokenBuffer(
			llm,
			10000,
		)
		mutex.Lock()
		mp[sessionID] = buf
		mutex.Unlock()
	}

	return buf, ok, nil
}

func (d *AIHelperDomain) RetrieveRelevantDocs(ctx context.Context, title, question string, scoreThreshold float32, topK int) ([]schema.Document, error) {
	docs, err := d.Qdrant.SearchSimilarDocuments(ctx, title, question, dao.WithScoreThreshold(scoreThreshold), dao.WithTopK(topK))
	if err != nil {
		return nil, fmt.Errorf("检索相关文档失败: %w", err)
	}

	return docs, nil
}

func (d *AIHelperDomain) SaveHistory(ctx context.Context, question, answer string, session *ChatSession) error {
	// 1. 会话内容
	if err := d.HistoryRepo.CreateHistory(ctx, &model.History{
		SessionID: session.SessionID,
		Question:  question,
		Answer:    answer,
	}); err != nil {
		return fmt.Errorf("保存历史记录失败: %w", err)
	}

	// 2. 会话 Session
	if session.IsNew {
		err := d.HistorySessionRepo.CreateHistorySession(ctx, &model.HistorySession{
			UserID:    session.UserID,
			SessionID: session.SessionID,
			Title:     question,
		})
		if err != nil {
			return fmt.Errorf("创建新会话失败: %w", err)
		}

		// 避免重复存储
		session.IsNew = false
	}

	// 3. 缓存
	if err := session.MemoryBuf.SaveContext(ctx, map[string]any{"question": question}, map[string]any{"answer": answer}); err != nil {
		return fmt.Errorf("保存历史记录失败: %w", err)
	}

	return nil
}

func (d *AIHelperDomain) BuildContext(ctx context.Context, buf *memory.ConversationTokenBuffer, req *types.AskQuestionRequest) ([]llms.MessageContent, error) {
	// 1. 检索相关文档
	docs, err := d.RetrieveRelevantDocs(ctx, req.Title, req.Question, req.ScoreThreshold, int(req.TopK))
	if err != nil {
		return nil, fmt.Errorf("检索相关文档失败: %w", err)
	}

	// 2. 获取历史记录
	history, err := buf.LoadMemoryVariables(context.Background(), map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("获取历史记录失败: %v", err)
	}

	// 3. 构建上下文
	var contextTexts []string
	for _, doc := range docs {
		contextTexts = append(contextTexts, doc.PageContent)
	}

	for _, h := range history {
		contextTexts = append(contextTexts, h.(string))
	}

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "请你仔细思考，读取上下文内容给出高质量的回答"),
		// 上下文
		llms.TextParts(llms.ChatMessageTypeHuman, strings.Join(contextTexts, "\n")),
		// 问题
		llms.TextParts(llms.ChatMessageTypeHuman, req.Question),
	}

	return content, nil
}

func (d *AIHelperDomain) BuildHistoryRespModel(ctx context.Context, histories []*model.History) []*types.GetChatHistoryResponse_ChatMessage {
	res := make([]*types.GetChatHistoryResponse_ChatMessage, 0, len(histories))
	for _, h := range histories {
		res = append(res, &types.GetChatHistoryResponse_ChatMessage{
			Question:   h.Question,
			Answer:     h.Answer,
			CreateTime: h.CreatedAt,
		})
	}

	return res
}
