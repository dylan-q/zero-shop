package logic

import (
	"context"
	"fmt"

	"zero-shop/app/user/rpc/internal/svc"
	"zero-shop/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *pb.GetUserByNameRequest) (*pb.GetUserInfoResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResponse{
		ID:       user.Id,
		UserName: user.Username,
	}, nil
}
