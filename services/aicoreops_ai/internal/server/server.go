package server

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/logic"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/types"
)

type AicoreopsAiServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedAIHelperServer
}

func NewAicoreopsAiServer(svcCtx *svc.ServiceContext) *AicoreopsAiServer {
	return &AicoreopsAiServer{
		svcCtx: svcCtx,
	}
}

// GetChatHistory 获取聊天历史
func (s *AicoreopsAiServer) GetChatHistory(ctx context.Context, req *types.GetChatHistoryRequest) (*types.GetChatHistoryResponse, error) {
	l := logic.NewAIHelperLogic(ctx, s.svcCtx)
	return l.GetChatHistory(req)
}

// UploadDocument 上传文档
func (s *AicoreopsAiServer) UploadDocument(ctx context.Context, req *types.UploadDocumentRequest) (*types.UploadDocumentResponse, error) {
	l := logic.NewAIHelperLogic(ctx, s.svcCtx)
	return l.UploadDocument(req)
}

// AskQuestion 实现 AI 助手的提问接口逻辑
func (s *AicoreopsAiServer) AskQuestion(stream types.AIHelper_AskQuestionServer) error {
	l := logic.NewAIHelperLogic(stream.Context(), s.svcCtx)
	return l.AskQuestion(stream)
}
