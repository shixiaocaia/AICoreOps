package logic

import (
	"aicoreops_ai/internal/domain"
	"aicoreops_ai/internal/svc"
	"aicoreops_ai/types"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/agents"
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
func (a *AIHelperLogic) AskQuestion(req *types.AskQuestionRequest) (*types.AskQuestionResponse, error) {
	// 检索相关文档
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
	fullPrompt := fmt.Sprintf("你是AICoreOps的AI助手，请你仔细思考，读取上下文内容给出高质量的回答:\n\n上下文:\n%s\n\n问题: %s",
		strings.Join(contextTexts, "\n"), req.Question)

	// 初始化并执行LLM调用
	executor, err := agents.Initialize(a.svcCtx.LLM, nil, agents.ConversationalReactDescription)
	if err != nil {
		a.Logger.Errorf("初始化AI代理失败: %v", err)
		return nil, fmt.Errorf("初始化AI代理失败: %v", err)
	}

	answer, err := chains.Run(a.ctx, executor, fullPrompt)
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
	// TODO: 实现获取对话历史的逻辑
	return &types.GetChatHistoryResponse{}, nil
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

	return &types.UploadDocumentResponse{}, nil
}
