package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"zero-shop/app/pay/rpc/internal/svc"
	"zero-shop/app/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetThirdByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetThirdByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetThirdByOrderSnLogic {
	return &GetThirdByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetThirdByOrderSnLogic) GetThirdByOrderSn(in *pb.GetThirdByOrderSnReq) (*pb.ThirdResp, error) {
	// todo: add your logic here and delete this line
	third, err := l.svcCtx.PayModel.FindOneByOrderSn(l.ctx, in.OrderSn)
	if err != nil {
		return nil, err
	}
	var thirdResp pb.ThirdResp
	_ = copier.Copy(&thirdResp, third)
	return &thirdResp, nil
}
