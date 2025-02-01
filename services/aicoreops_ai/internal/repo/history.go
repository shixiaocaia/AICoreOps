package repo

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_ai/internal/model"
)

type HistoryRepo interface {
	CreateHistory(ctx context.Context, history *model.History) error
	GetHistoryBySessionID(ctx context.Context, sessionID string) ([]*model.History, error)
	GetHistoryList(ctx context.Context, userId int64, offset, limit int) ([]*model.History, error)
	UpdateHistory(ctx context.Context, history *model.History) error
	DeleteHistory(ctx context.Context, id int64) error
}
