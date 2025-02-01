package logic

import (
	"context"
	"fmt"
	"io"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/domain"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/types"
	"github.com/google/uuid"

	"github.com/tmc/langchaingo/llms"
	"github.com/zeromicro/go-zero/core/logx"
)

type AIHelperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	domain *domain.AIHelperDomain
}

func NewAIHelperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AIHelperLogic {
	return &AIHelperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		domain: domain.NewAIHelperDomain(svcCtx.DB, svcCtx.Qdrant),
	}
}

// GetHistoryList 获取历史会话列表
func (a *AIHelperLogic) GetHistoryList(req *types.GetHistoryListRequest) (*types.GetHistoryListResponse, error) {
	uid := a.ctx.Value("userId").(int64)
	if uid == 0 {
		a.Logger.Error("[获取历史会话列表失败]: 用户ID为空")
		return nil, fmt.Errorf("用户ID为空")
	}

	histories, err := a.domain.GetHistorySessionList(a.ctx, uid, int(req.PageSize), int(req.Page))
	if err != nil {
		a.Logger.Errorf("[获取历史会话列表失败]: %v", err)
		return nil, fmt.Errorf("获取历史会话列表失败: %v", err)
	}

	return &types.GetHistoryListResponse{
		Code:    0,
		Message: "success",
		Data:    histories,
	}, nil
}

// GetChatHistory 获取指定 sessionID 的历史聊天记录
func (a *AIHelperLogic) GetChatHistory(req *types.GetChatHistoryRequest) (*types.GetChatHistoryResponse, error) {
	// 1. 获取历史记录
	histories, err := a.domain.GetHistoryBySessionID(a.ctx, req.SessionId)
	if err != nil {
		a.Logger.Errorf("获取历史记录失败: %v", err)
		return nil, fmt.Errorf("获取历史记录失败: %v", err)
	}

	// 2. 异步加载记录到 memoryBuf
	if err = a.domain.LoadHistoryToMemory(a.ctx, a.svcCtx, req.SessionId, histories); err != nil {
		a.Logger.Errorf("加载历史记录失败: %v", err)
		return nil, fmt.Errorf("加载历史记录失败: %v", err)
	}

	return &types.GetChatHistoryResponse{
		Code:    0,
		Message: "success",
		Data:    &types.GetChatHistoryResponse_ChatHistoryData{Messages: histories, Total: int32(len(histories))},
	}, nil
}

// UploadDocument 上传运维文档
func (a *AIHelperLogic) UploadDocument(req *types.UploadDocumentRequest) (*types.UploadDocumentResponse, error) {
	if err := a.domain.UploadDocument(a.ctx, req.Title, req.Content); err != nil {
		a.Logger.Errorf("上传文档失败: %v", err)
		return nil, fmt.Errorf("上传文档失败: %v", err)
	}

	return &types.UploadDocumentResponse{
		Code:    0,
		Message: "success",
	}, nil
}

// CreateNewChat 创建新聊天
func (a *AIHelperLogic) CreateNewChat(req *types.CreateNewChatRequest) (*types.CreateNewChatResponse, error) {
	// 1. 生成sessionID
	sessionID := uuid.New().String()

	// TODO 2. 只有发起对话才会真正创建会话
	// _, err := a.svcCtx.HistorySessionModel.Insert(a.ctx, &model.HistorySession{
	// 	UserId:    req.UserId,
	// 	SessionId: sessionID,
	// 	CreatedAt: time.Now(),
	// })
	// if err != nil {
	// 	a.Logger.Errorf("创建新会话失败: %v", err)
	// 	return nil, fmt.Errorf("创建新会话失败: %v", err)
	// }

	return &types.CreateNewChatResponse{
		Code:    0,
		Message: "success",
		Data: &types.CreateNewChatResponse_SessionData{
			SessionId: sessionID,
		},
	}, nil
}

// AskQuestion 实现 AI 助手的提问接口逻辑，使用双向流式 RPC
func (a *AIHelperLogic) AskQuestion(stream types.AIHelper_AskQuestionServer) error {
	// 1. check session
	userID, sessionID, err := a.domain.CheckSession(a.ctx)
	if err != nil {
		a.Logger.Errorf("检查会话失败: %v", err)
		return fmt.Errorf("检查会话失败: %v", err)
	}

	// 2. get memoryBuf
	buf, newSession, err := a.domain.GetMemoryBuf(a.ctx, sessionID, a.svcCtx.LLM, a.svcCtx.MemoryBuf, a.svcCtx.Mutex)
	if err != nil {
		a.Logger.Errorf("获取 MemoryBuf 失败: %v", err)
		return fmt.Errorf("获取 MemoryBuf 失败: %v", err)
	}

	// 3. Ask & Reply
	for {
		// 3.1从流中接收请求
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				a.Logger.Info("释放流: %v", sessionID)
				return nil
			}
			a.Logger.Errorf("流: %v, 接收请求失败: %v", sessionID, err)
			return fmt.Errorf("流: %v, 接收请求失败: %v", sessionID, err)
		}

		a.Logger.Infof("成功接收请求: %v", req.Question)

		// 3.2 构建上下文
		content, err := a.domain.BuildContext(a.ctx, buf, req)
		if err != nil {
			a.Logger.Errorf("构建上下文失败: %v", err)
			return fmt.Errorf("构建上下文失败: %v", err)
		}

		// 3.3 生成回答 & 流式发送
		completion, err := a.svcCtx.LLM.GenerateContent(a.ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			err := stream.Send(&types.AskQuestionResponse{
				Code:    0,
				Message: "success",
				Data:    &types.AskQuestionResponse_AnswerData{Answer: string(chunk), SessionId: sessionID},
			})
			if err != nil {
				a.Logger.Errorf("发送响应失败: %v", err)
				return fmt.Errorf("发送响应失败: %v", err)
			}
			return nil
		}))
		if err != nil {
			a.Logger.Errorf("生成回答失败: %v", err)
			return fmt.Errorf("生成回答失败: %v", err)
		}

		// 3.4 保存对话历史
		err = a.domain.SaveHistory(a.ctx, completion.Choices[0].Content, req, newSession, userID, sessionID, buf)
		if err != nil {
			a.Logger.Errorf("保存对话历史失败: %v", err)
			return fmt.Errorf("保存对话历史失败: %v", err)
		}

		a.Logger.Infof("成功生成对话: %v", sessionID)
	}
}
