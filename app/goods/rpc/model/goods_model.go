package model

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsModel = (*customGoodsModel)(nil)

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
		FindCount(ctx context.Context) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64, orderBy string) ([]*Goods, error)
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn sqlx.SqlConn) GoodsModel {
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn),
	}
}
func (m *defaultGoodsModel) FindPageListByPage(ctx context.Context, page, pageSize int64, orderBy string) ([]*Goods, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var resp []*Goods
	query, values, err := sq.Select("*").From(m.table).OrderBy("id DESC").Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *defaultGoodsModel) FindCount(ctx context.Context) (int64, error) {

	var resp int64
	query := fmt.Sprintf("select count(*) from %s", m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
