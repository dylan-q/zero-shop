package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"zero-shop/app/pay/rpc/internal/svc"
	"zero-shop/app/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThirdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdLogic {
	return &ThirdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThirdLogic) Third(in *pb.ThirdReq) (*pb.ThirdResp, error) {
	// todo: add your logic here and delete this line
	third, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var thirdResp pb.ThirdResp
	_ = copier.Copy(&thirdResp, third)
	return &thirdResp, nil
}
