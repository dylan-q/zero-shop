package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-shop/app/pay/rpc/pb"
)

var _ ThirdPayModel = (*customThirdPayModel)(nil)

type (
	// ThirdPayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPayModel.
	ThirdPayModel interface {
		thirdPayModel
		FindOneByOrderSn(ctx context.Context, sn string) (*ThirdPay, error)
		UpdateByOrderSn(ctx context.Context, req *pb.UpdateByOrderSnReq) error
	}

	customThirdPayModel struct {
		*defaultThirdPayModel
	}
)

// NewThirdPayModel returns a model for the database table.
func NewThirdPayModel(conn sqlx.SqlConn) ThirdPayModel {
	return &customThirdPayModel{
		defaultThirdPayModel: newThirdPayModel(conn),
	}
}

func (m *defaultThirdPayModel) FindOneByOrderSn(ctx context.Context, sn string) (*ThirdPay, error) {
	query := fmt.Sprintf("select %s from %s where `order_sn` = ? limit 1", thirdPayRows, m.table)
	var resp ThirdPay
	err := m.conn.QueryRowCtx(ctx, &resp, query, sn)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultThirdPayModel) UpdateByOrderSn(ctx context.Context, req *pb.UpdateByOrderSnReq) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, thirdPayRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, req.Sn, req.PayMode, req.TradeType, req.TradeState, req.PayTotal, req.TransactionId, req.TradeStateDesc, req.PayStatus, req.PayTime)
	return err

}
