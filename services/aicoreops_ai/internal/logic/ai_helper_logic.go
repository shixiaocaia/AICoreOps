package logic

import (
	"aicoreops_ai/internal/domain"
	"aicoreops_ai/internal/model"
	"aicoreops_ai/internal/svc"
	"aicoreops_ai/types"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/tmc/langchaingo/llms"
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

// AskQuestion 实现 AI 助手的提问接口逻辑，使用双向流式 RPC
func (a *AIHelperLogic) AskQuestion(stream types.AIHelper_AskQuestionServer) error {
	for {
		// 从流中接收请求
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				a.Logger.Info("接收请求完成")
				return nil
			}
			a.Logger.Errorf("接收请求失败: %v", err)
			return fmt.Errorf("接收请求失败: %v", err)
		}

		docRetrieved, err := domain.RetrieveRelevantDocs(a.svcCtx.Qdrant, req.Question)
		if err != nil {
			a.Logger.Errorf("检索相关文档失败: %v", err)
			return fmt.Errorf("检索相关文档失败: %v", err)
		}

		// 构建上下文
		var contextTexts []string
		for _, doc := range docRetrieved {
			contextTexts = append(contextTexts, doc.PageContent)
		}

		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "你是AICoreOps的AI助手，请你仔细思考，读取上下文内容给出高质量的回答"),
			// 上下文
			llms.TextParts(llms.ChatMessageTypeHuman, strings.Join(contextTexts, "\n")),
			// 问题
			llms.TextParts(llms.ChatMessageTypeHuman, req.Question),
		}
		completion, err := a.svcCtx.LLM.GenerateContent(a.ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// 将生成的内容流式发送回客户端
			err := stream.Send(&types.AskQuestionResponse{
				Code:    0,
				Message: "success",
				Data:    &types.AskQuestionResponse_AnswerData{Answer: string(chunk)},
			})
			if err != nil {
				a.Logger.Errorf("发送响应失败: %v", err)
				return fmt.Errorf("发送响应失败: %v", err)
			}
			return nil
		}))

		if err != nil {
			log.Fatal(err)
		}
		_ = completion
	}
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
