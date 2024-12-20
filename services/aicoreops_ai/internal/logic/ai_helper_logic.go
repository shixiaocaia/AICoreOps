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
	"github.com/tmc/langchaingo/memory"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
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
	// 从元数据中获取 sessionID, 加载历史记录
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		a.Logger.Error("无法从上下文中获取元数据")
		return fmt.Errorf("无法从上下文中获取元数据")
	}

	sessionIDs := md["sessionid"]
	if len(sessionIDs) == 0 {
		a.Logger.Error("sessionID 未设置或为空")
		return fmt.Errorf("sessionID 未设置或为空")
	}
	sessionID := sessionIDs[0]

	buf, ok := a.svcCtx.MemoryBuf[sessionID]
	if !ok {
		buf = memory.NewConversationTokenBuffer(
			a.svcCtx.LLM,
			10000,
		)
		a.svcCtx.MemoryBuf[sessionID] = buf
	}

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
		// 获取历史记录
		history, err := buf.LoadMemoryVariables(context.Background(), map[string]any{})
		if err != nil {
			log.Println("Failed to load memory variables:", err)
			return fmt.Errorf("获取历史记录失败: %v", err)
		}

		// 检索相关文档
		docRetrieved, err := domain.RetrieveRelevantDocs(a.svcCtx.Qdrant, req.Question)
		if err != nil {
			a.Logger.Errorf("检索相关文档失败: %v", err)
			return fmt.Errorf("检索相关文档失败: %v", err)
		}

		// 构建上下文
		var contextTexts []string
		for _, doc := range docRetrieved {
			contextTexts = append(contextTexts, doc.PageContent)
			a.Logger.Infof("文档: %v", doc.PageContent)
		}

		for _, h := range history {
			contextTexts = append(contextTexts, h.(string))
			a.Logger.Infof("历史记录: %v", h.(string))
		}

		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "你是AICoreOps的AI助手，请你仔细思考，读取上下文内容给出高质量的回答"),
			// 上下文
			llms.TextParts(llms.ChatMessageTypeHuman, strings.Join(contextTexts, "\n")),
			// 问题
			llms.TextParts(llms.ChatMessageTypeHuman, req.Question),
		}

		// 生成回答，流式发送回客户端
		completion, err := a.svcCtx.LLM.GenerateContent(a.ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
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
			return fmt.Errorf("生成回答失败: %v", err)
		}

		// 持久化历史记录
		historyModel := model.NewHistoryModel(a.svcCtx.DB)
		err = historyModel.Create(&model.History{
			SessionID: req.SessionId,
			Question:  req.Question,
			Answer:    completion.Choices[0].Content,
		})

		if err != nil {
			a.Logger.Errorf("持久化历史记录失败: %v", err)
		}

		// 保存历史记录到 memoryBuf
		err = buf.SaveContext(a.ctx, map[string]any{"question": req.Question}, map[string]any{"answer": completion.Choices[0].Content})
		if err != nil {
			a.Logger.Errorf("保存历史记录失败: %v", err)
		}
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

	// 异步加载记录到 memoryBuf
	// TODO 加载精简的一条关键信息，而不是全部
	go func() {
		if _, ok := a.svcCtx.MemoryBuf[req.SessionId]; ok {
			a.Logger.Infof("历史记录已加载: %v", req.SessionId)
			return
		}
		buf := memory.NewConversationTokenBuffer(
			a.svcCtx.LLM,
			10000,
		)
		for _, h := range histories {
			err := buf.SaveContext(a.ctx, map[string]any{"question": h.Question}, map[string]any{"answer": h.Answer})
			if err != nil {
				a.Logger.Errorf("保存历史记录失败: %v", err)
			}
		}
		a.svcCtx.MemoryBuf[req.SessionId] = buf
		a.Logger.Infof("加载历史记录成功: %v", req.SessionId)
	}()

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
