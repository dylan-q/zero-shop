package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"zero-shop/app/order/rpc/internal/svc"
	"zero-shop/app/order/rpc/pb"

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

func (l *ListLogic) List(in *pb.ListReq) (*pb.ListResp, error) {
	// todo: add your logic here and delete this line
	total, err := l.svcCtx.OrderModel.FindCount(l.ctx)
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.OrderModel.FindPageListByPage(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	var infoList []*pb.OrderInfoResp
	for _, v := range result {
		var info pb.OrderInfoResp
		_ = copier.Copy(&info, v)
		infoList = append(infoList, &info)
	}
	return &pb.ListResp{
		Total: total,
		Data:  infoList,
	}, nil
}
