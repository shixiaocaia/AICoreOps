package logic

import (
	"aicoreops_ai/internal/domain"
	"aicoreops_ai/internal/model"
	"aicoreops_ai/internal/svc"
	"aicoreops_ai/types"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/zeromicro/go-zero/core/logx"
)

type AIHelperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAIHelperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AIHelperLogic {
	return &AIHelperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AskQuestion 实现 AI 助手的提问接口逻辑
// TODO 1. 保存历史对话 2. 流式对话
func (a *AIHelperLogic) AskQuestion(req *types.AskQuestionRequest) (*types.AskQuestionResponse, error) {
	// 查询历史文档，并构建上下文存入 memoryBuf

	// 升级为流式对话

	// 检索相关文档
	// 每一次问题都检索知识库吗
	docRetrieved, err := domain.RetrieveRelevantDocs(a.svcCtx.Qdrant, req.Question)
	if err != nil {
		a.Logger.Errorf("检索相关文档失败: %v", err)
		return nil, fmt.Errorf("检索相关文档失败: %v", err)
	}

	// 构建上下文
	var contextTexts []string
	for _, doc := range docRetrieved {
		contextTexts = append(contextTexts, doc.PageContent)
	}

	// 构建完整提示词
	// TODO 检索用户设置的提示词
	fullPrompt := fmt.Sprintf("你是AICoreOps的AI助手，请你仔细思考，读取上下文内容给出高质量的回答:\n\n上下文:\n%s\n\n问题: %s",
		strings.Join(contextTexts, "\n"), req.Question)

	answer, err := chains.Run(a.ctx, a.svcCtx.Executor, fullPrompt)
	if err != nil {
		a.Logger.Errorf("生成回答失败: %v", err)
		return nil, fmt.Errorf("生成回答失败: %v", err)
	}

	return &types.AskQuestionResponse{
		Code:    0,
		Message: "success",
		Data:    &types.AskQuestionResponse_AnswerData{Answer: answer},
	}, nil
}

// GetChatHistory 获取用户的对话历史记录
func (a *AIHelperLogic) GetChatHistory(req *types.GetChatHistoryRequest) (*types.GetChatHistoryResponse, error) {
	historyModel := model.NewHistoryModel(a.svcCtx.DB)

	histories, err := historyModel.GetBySessionID(req.SessionId)
	if err != nil {
		a.Logger.Errorf("查询历史记录失败: %v", err)
		return nil, fmt.Errorf("查询历史记录失败: %v", err)
	}

	// 构建返回历史
	res := make([]*types.GetChatHistoryResponse_ChatMessage, 0, len(histories))
	for _, h := range histories {
		res = append(res, &types.GetChatHistoryResponse_ChatMessage{
			Question:   h.Question,
			Answer:     h.Answer,
			CreateTime: h.CreatedAt.Unix(),
		})
	}

	return &types.GetChatHistoryResponse{
		Code:    0,
		Message: "success",
		Data:    &types.GetChatHistoryResponse_ChatHistoryData{Messages: res, Total: int32(len(histories))},
	}, nil
}

// UploadDocument 上传运维文档，丰富 AI 助手的知识库
func (a *AIHelperLogic) UploadDocument(req *types.UploadDocumentRequest) (*types.UploadDocumentResponse, error) {
	// 解码 Base64 内容
	fileBytes, err := base64.StdEncoding.DecodeString(req.Content)
	if err != nil {
		a.Logger.Errorf("解码 Base64 内容失败: %v", err)
		return nil, fmt.Errorf("解码 Base64 内容失败: %v", err)
	}

	// TODO 存储 title

	// 将文本文件拆分成文档块
	docs, err := domain.TextToChunks(fileBytes)
	if err != nil {
		a.Logger.Errorf("拆分文本失败: %v", err)
		return nil, fmt.Errorf("拆分文本失败: %v", err)
	}

	// 将文档存储到 Qdrant 向量存储中
	err = domain.InsertDocsToQdrantStore(a.svcCtx.Qdrant, docs)
	if err != nil {
		a.Logger.Errorf("添加文档失败: %v", err)
		return nil, fmt.Errorf("添加文档失败: %v", err)
	}

	return &types.UploadDocumentResponse{
		Code:    0,
		Message: "success",
	}, nil
}
