package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-shop/app/pay/api/internal/svc"
	"zero-shop/app/pay/api/internal/types"
)

type NotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyLogic {
	return &NotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyLogic) Notify(req *types.NotifyReq) (resp *types.NotifyResp, err error) {
	// todo: add your logic here and delete this line
	//pay, err := l.svcCtx.Config.PayRpcConf.BuildTarget()
	//if err != nil {
	//	return nil, status.Error(codes.Internal, "回调失败")
	//}
	//gid := dtmgrpc.MustGenGid(l.svcCtx.Config.MicroService.Target)
	//var client, err = alipay.New("", "", true)
	//var noti, err = client.DecodeNotification(req)
	return
}
