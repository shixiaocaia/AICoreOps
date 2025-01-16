package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistorySessionModel = (*customHistorySessionModel)(nil)

type (
	// HistorySessionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistorySessionModel.
	HistorySessionModel interface {
		historySessionModel
		withSession(session sqlx.Session) HistorySessionModel
		FindAll(ctx context.Context, userId string) ([]*HistorySession, error)
	}

	customHistorySessionModel struct {
		*defaultHistorySessionModel
	}
)

// NewHistorySessionModel returns a model for the database table.
func NewHistorySessionModel(conn sqlx.SqlConn) HistorySessionModel {
	return &customHistorySessionModel{
		defaultHistorySessionModel: newHistorySessionModel(conn),
	}
}

func (m *customHistorySessionModel) withSession(session sqlx.Session) HistorySessionModel {
	return NewHistorySessionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customHistorySessionModel) FindAll(ctx context.Context, userId string) ([]*HistorySession, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? ORDER BY created_at DESC", historySessionRows, m.table)
	var resp []*HistorySession
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	return resp, err
}
