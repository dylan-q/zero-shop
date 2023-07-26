package logic

import (
	"context"
	"fmt"

	//"github.com/dtm-labs/dtmgrpc"
	"github.com/dtm-labs/client/dtmgrpc"
	"zero-shop/app/order/rpc/pb"
	paypb "zero-shop/app/pay/rpc/pb"
	"zero-shop/pkg/utils"

	"zero-shop/app/pay/api/internal/svc"
	"zero-shop/app/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateThirdReq) (resp *types.CreateThirdResp, err error) {
	// todo: add your logic here and delete this line
	orderRpcBuild, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		return nil, err
	}
	payRpcBuild, err := l.svcCtx.Config.PayRpcConf.BuildTarget()
	if err != nil {
		return nil, err
	}
	sn := utils.GenSn("ORD-")
	gid := dtmgrpc.MustGenGid(l.svcCtx.Config.MicroService.Target)
	fmt.Println(gid)
	//fmt.Println(payRpcBuild)
	msg := dtmgrpc.NewMsgGrpc(l.svcCtx.Config.MicroService.Target, gid).Add(orderRpcBuild+"/pb.order/Create", &pb.CreateOrderReq{
		UserId:      1,
		GoodsId:     req.GoodsId,
		MarketPrice: req.MarketPrice,
		SalePrice:   req.SalePrice,
		OrderStatus: 0,
		OrderSn:     sn,
	}).Add(payRpcBuild+"/pb.pay/Create", &paypb.PayReq{
		OrderSn: sn,
		UserId:  1,
	})
	err = msg.Submit()
	//fmt.Println("eee" + err.Error())
	if err != nil {
		return nil, err
	}

	return &types.CreateThirdResp{OrderSn: sn}, nil
}
