package user

import (
	"context"
	"github.com/pkg/errors"
	rpc "zero-shop/app/user/rpc/user"
	my_utils "zero-shop/pkg/utils"

	"zero-shop/app/user/api/internal/svc"
	"zero-shop/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserDetailReq) (resp *types.UserDetailRsq, err error) {
	// todo: add your logic here and delete this line
	uid := my_utils.GetUidFromCtx(l.ctx)
	if uid <= 0 {
		return nil, errors.New("解析token失败")
	}
	user, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &rpc.GetUserInfoRequest{
		ID: uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserDetailRsq{
		ID:       uid,
		UserName: user.UserName,
	}, nil
}
