package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindCount(ctx context.Context) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64) ([]*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn),
	}
}
func (m *defaultOrderModel) FindPageListByPage(ctx context.Context, page, pageSize int64) ([]*Order, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var resp []*Order
	query, values, err := sq.Select("*").From(m.table).OrderBy("id DESC").Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *defaultOrderModel) FindCount(ctx context.Context) (int64, error) {

	var resp int64
	query, _, err := sq.Select("COUNT(*)").From(m.table).ToSql()
	err = m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
