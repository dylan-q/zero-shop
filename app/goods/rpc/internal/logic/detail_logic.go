package logic

import (
	"context"
	"github.com/jinzhu/copier"

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
	goods, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var info pb.GoodsInfoResponse
	_ = copier.Copy(&info, goods)
	category, err := l.svcCtx.CateModel.FindOne(l.ctx, goods.CategoryId)
	var cate pb.CategoryInfoResponse
	if err != nil {
		cate.Name = ""
		cate.Id = 0
	} else {
		cate.Name = category.Name
		cate.Id = category.Id
	}
	info.Category = &cate
	return &info, nil
}
