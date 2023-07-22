package logic

import (
	"context"
	"zero-shop/app/order/rpc/model"
	"zero-shop/pkg/utils"

	"zero-shop/app/order/rpc/internal/svc"
	"zero-shop/app/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	// todo: add your logic here and delete this line
	insert := new(model.Order)
	sn := utils.GenSn("ORD")
	insert.UserId = in.UserId
	insert.GoodsId = in.GoodsId
	insert.MarketPrice = in.MarketPrice
	insert.SalePrice = in.SalePrice
	insert.OrderStatus = int64(0)
	insert.Sn = sn
	res, err := l.svcCtx.OrderModel.Insert(l.ctx, insert)
	if err != nil {
		return nil, err
	}
	oid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResp{
		Id:      oid,
		Sn:      sn,
		GoodsId: in.GoodsId,
		UserId:  in.UserId,
	}, nil
}
