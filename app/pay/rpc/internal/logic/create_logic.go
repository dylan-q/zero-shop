package logic

import (
	"context"
	"fmt"
	"time"
	"zero-shop/app/pay/rpc/model"

	"zero-shop/app/pay/rpc/internal/svc"
	"zero-shop/app/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.PayReq) (*pb.PayResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入到pay-rpc")
	var data model.ThirdPay
	data.UserId = in.UserId
	data.OrderSn = in.OrderSn
	data.PayStatus = 0
	data.CreateTime = time.Now()
	insert, err := l.svcCtx.PayModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.PayResp{
		Id: id,
	}, nil
}
