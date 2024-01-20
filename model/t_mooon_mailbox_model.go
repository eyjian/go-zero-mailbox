package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TMooonMailboxModel = (*customTMooonMailboxModel)(nil)

type (
	// TMooonMailboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTMooonMailboxModel.
	TMooonMailboxModel interface {
		tMooonMailboxModel
		withSession(session sqlx.Session) TMooonMailboxModel
	}

	customTMooonMailboxModel struct {
		*defaultTMooonMailboxModel
	}
)

// NewTMooonMailboxModel returns a model for the database table.
func NewTMooonMailboxModel(conn sqlx.SqlConn) TMooonMailboxModel {
	return &customTMooonMailboxModel{
		defaultTMooonMailboxModel: newTMooonMailboxModel(conn),
	}
}

func (m *customTMooonMailboxModel) withSession(session sqlx.Session) TMooonMailboxModel {
	return NewTMooonMailboxModel(sqlx.NewSqlConnFromSession(session))
}
