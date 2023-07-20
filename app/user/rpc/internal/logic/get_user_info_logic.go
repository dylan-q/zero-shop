package logic

import (
	"context"

	"zero-shop/app/user/rpc/internal/svc"
	"zero-shop/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.ID)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResponse{
		ID:       user.Id,
		UserName: user.Username,
	}, nil
}
