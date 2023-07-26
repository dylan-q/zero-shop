package logic

import (
	"context"

	"zero-shop/app/order/rpc/internal/svc"
	"zero-shop/app/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.OrderModel.UpdateOrderStatus(l.ctx, &pb.UpdateOrderStatusReq{
		OrderId:     in.OrderId,
		OrderStatus: in.OrderStatus,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateOrderStatusResp{
		OrderId: in.OrderId,
	}, nil
}
