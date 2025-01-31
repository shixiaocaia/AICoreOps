package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/model"
)

type HistorySessionRepo interface {
	CreateHistorySession(ctx context.Context, session *model.HistorySession) error
	GetHistorySessionByID(ctx context.Context, id int64) (*model.HistorySession, error)
	GetHistorySessionList(ctx context.Context, userId int64, offset, limit int) ([]*model.HistorySession, error)
	UpdateHistorySession(ctx context.Context, session *model.HistorySession) error
	DeleteHistorySession(ctx context.Context, sessionId int64) error
}
