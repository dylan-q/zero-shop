package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-shop/app/goods/rpc/pb"

	"zero-shop/app/goods/api/internal/svc"
	"zero-shop/app/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailReq) (resp *types.GoodsInfo, err error) {
	// todo: add your logic here and delete this line
	goods, err := l.svcCtx.GoodsRpc.Detail(l.ctx, &pb.DetailRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var result types.GoodsInfo
	_ = copier.Copy(&result, goods)
	return &result, nil
}
