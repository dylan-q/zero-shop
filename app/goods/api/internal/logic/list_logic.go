package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-shop/app/goods/rpc/pb"

	"zero-shop/app/goods/api/internal/svc"
	"zero-shop/app/goods/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (resp *types.ListResp, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.GoodsRpc.List(l.ctx, &pb.ListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var resList []types.GoodsInfo

	if len(list.Data) > 0 {
		for _, v := range list.Data {
			var info types.GoodsInfo
			_ = copier.Copy(&info, v)
			resList = append(resList, info)
		}
	}

	return &types.ListResp{
		Data:  resList,
		Total: list.Total,
	}, nil
}
