package logic

import (
	"context"

	"zero-shop/app/goods/rpc/internal/svc"
	"zero-shop/app/goods/rpc/pb"

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

func (l *DetailLogic) Detail(in *pb.DetailRequest) (*pb.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodsInfoResponse{}, nil
}
