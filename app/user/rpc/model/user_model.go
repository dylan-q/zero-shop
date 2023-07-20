package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByNameAndPassword(ctx context.Context, UserName string, password string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
func (m *defaultUserModel) FindOneByNameAndPassword(ctx context.Context, username string, password string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? and `password` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, username, password)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
