package logic

import (
	"context"
	"zero-shop/app/user/rpc/model"
	s_utils "zero-shop/pkg/utils"

	"zero-shop/app/user/rpc/internal/svc"
	"zero-shop/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	// todo: add your logic here and delete this line
	user := &model.User{Username: in.Username, Password: s_utils.MD5(in.Password)}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	if id <= 0 {
		return nil, err
	}
	return &pb.UserRegisterResponse{
		ID:       id,
		UserName: in.Username,
	}, nil
}
