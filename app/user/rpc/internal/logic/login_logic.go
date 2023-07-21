package logic

import (
	"context"
	"fmt"
	s_utils "zero-shop/pkg/utils"

	"zero-shop/app/user/rpc/internal/svc"
	"zero-shop/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in)
	user, err := l.svcCtx.UserModel.FindOneByNameAndPassword(l.ctx, in.Username, s_utils.MD5(in.Password))
	if err != nil {
		return nil, err
	}
	return &pb.UserLoginResponse{
		ID:       user.Id,
		UserName: user.Username,
	}, nil
}
