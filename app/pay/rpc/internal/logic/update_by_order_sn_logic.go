package logic

import (
	"context"
	"zero-shop/app/pay/rpc/internal/svc"
	"zero-shop/app/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateByOrderSnLogic {
	return &UpdateByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateByOrderSnLogic) UpdateByOrderSn(in *pb.UpdateByOrderSnReq) (*pb.UpdateByOrderSnResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.PayModel.UpdateByOrderSn(l.ctx, &pb.UpdateByOrderSnReq{
		Sn:             in.Sn,
		UpdateTime:     in.UpdateTime,
		PayMode:        in.PayMode,
		TradeType:      in.TradeType,
		TradeState:     in.TradeState,
		PayTotal:       in.PayTotal,
		TransactionId:  in.TransactionId,
		TradeStateDesc: in.TradeStateDesc,
		OrderSn:        in.OrderSn,
		PayStatus:      in.PayStatus,
		PayTime:        in.PayTime,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateByOrderSnResp{}, nil
}
