package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-shop/app/goods/rpc/goods"
	"zero-shop/app/pay/rpc/pay"

	"zero-shop/app/order/rpc/internal/svc"
	"zero-shop/app/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *pb.DetailReq) (*pb.OrderInfoResp, error) {
	// todo: add your logic here and delete this line
	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	orderGoods, err := l.svcCtx.GoodsRpc.Detail(l.ctx, &goods.DetailRequest{
		Id: order.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	var goodsInfo pb.GoodsInfo
	_ = copier.Copy(&goodsInfo, orderGoods)
	thirdInfo, err := l.svcCtx.PayRpc.GetThirdByOrderSn(l.ctx, &pay.GetThirdByOrderSnReq{
		OrderSn: order.Sn,
	})
	var third pb.ThirdInfo
	_ = copier.Copy(&third, thirdInfo)
	return &pb.OrderInfoResp{
		Id:          order.Id,
		UserId:      order.UserId,
		OrderStatus: order.OrderStatus,
		GoodsId:     order.GoodsId,
		Goods:       &goodsInfo,
		MarketPrice: order.MarketPrice,
		SalePrice:   order.SalePrice,
		Sn:          order.Sn,
		CreateTime:  order.CreateTime.Unix(),
		Third:       &third,
	}, nil
}
