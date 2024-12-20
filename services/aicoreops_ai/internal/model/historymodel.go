package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistoryModel = (*customHistoryModel)(nil)

type (
	// HistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistoryModel.
	HistoryModel interface {
		historyModel
		withSession(session sqlx.Session) HistoryModel
		FindAll(ctx context.Context, sessionId string) ([]*History, error)
	}

	customHistoryModel struct {
		*defaultHistoryModel
	}
)

// NewHistoryModel returns a model for the database table.
func NewHistoryModel(conn sqlx.SqlConn) HistoryModel {
	return &customHistoryModel{
		defaultHistoryModel: newHistoryModel(conn),
	}
}

func (m *customHistoryModel) withSession(session sqlx.Session) HistoryModel {
	return NewHistoryModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customHistoryModel) FindAll(ctx context.Context, sessionId string) ([]*History, error) {
	var histories []*History
	err := m.conn.QueryRows(&histories, "SELECT * FROM history WHERE session_id = ?", sessionId)
	if err != nil {
		return nil, err
	}
	return histories, nil
}
