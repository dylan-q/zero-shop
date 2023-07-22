package logic

import (
	"context"

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

	return &pb.OrderInfoResp{}, nil
}
