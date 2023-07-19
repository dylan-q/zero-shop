package user

import (
	"context"
	"strconv"
	"time"
	"zero-shop/app/user/rpc/user"
	s_utils "zero-shop/pkg/utils"

	"zero-shop/app/user/api/internal/svc"
	"zero-shop/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	// todo: add your logic here and delete this line
	loginUser, err := l.svcCtx.UserRpc.Login(l.ctx, &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	if loginUser == nil {
		return nil, err
	}
	token, err := s_utils.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, strconv.FormatInt(loginUser.ID, 10))
	if err != nil {
		return nil, err
	}
	return &types.LoginRsp{
		Token: token,
	}, nil
}
