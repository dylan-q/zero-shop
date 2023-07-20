package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
	"zero-shop/app/user/api/internal/svc"
	"zero-shop/app/user/api/internal/types"
	rpcuser "zero-shop/app/user/rpc/user"
	s_utils "zero-shop/pkg/utils"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUserByName(l.ctx, &rpcuser.GetUserByNameRequest{
		Username: req.Username,
	})
	if user != nil {
		return nil, errors.New("用户已存在")
	}
	result, err := l.svcCtx.UserRpc.Register(l.ctx, &rpcuser.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token, err := s_utils.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, strconv.FormatInt(result.ID, 10))
	if err != nil {
		return nil, err
	}
	return &types.LoginRsp{
		Token: token,
	}, nil
}
