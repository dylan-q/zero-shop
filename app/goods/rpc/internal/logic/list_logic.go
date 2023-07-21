package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"zero-shop/app/goods/rpc/internal/svc"
	"zero-shop/app/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pb.ListRequest) (*pb.ListResponse, error) {
	// todo: add your logic here and delete this line
	count, err := l.svcCtx.GoodsModel.FindCount(l.ctx)

	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.GoodsModel.FindPageListByPage(l.ctx, in.Page, in.PageSize, "id")
	if err != nil {
		return nil, err
	}
	var data []*pb.GoodsInfoResponse
	for _, v := range result {
		var cate pb.CategoryInfoResponse
		category, _ := l.svcCtx.CateModel.FindOne(l.ctx, v.CategoryId)
		_ = copier.Copy(&cate, category)
		var info pb.GoodsInfoResponse
		info.Id = v.Id
		info.GoodsName = v.GoodsName
		info.CategoryId = v.CategoryId
		info.ClickNum = v.ClickNum
		info.SoldNum = v.SoldNum
		info.OnSale = v.OnSale
		info.FavNum = v.FavNum
		info.CreateTime = v.CreateTime.Unix()
		info.GoodsBrief = v.GoodsBrief
		info.GoodsDetail = v.GoodsDetail.String
		info.IsHot = v.IsHot
		info.IsNew = v.IsNew
		info.MarketPrice = float32(v.MarketPrice)
		info.ShopPrice = float32(v.ShopPrice)
		info.ShipFree = v.ShipFree
		info.Category = &cate
		data = append(data, &info)
	}
	return &pb.ListResponse{
		Total: count,
		Data:  data,
	}, nil
}
