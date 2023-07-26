package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/app/goods/rpc/goods"
	"zero-shop/app/order/rpc/internal/config"
	"zero-shop/app/order/rpc/model"
	"zero-shop/app/pay/rpc/pay"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	GoodsRpc   goods.Goods
	PayRpc     pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(sqlx.NewMysql(c.DB.DataSource)),
		GoodsRpc:   goods.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
		PayRpc:     pay.NewPay(zrpc.MustNewClient(c.PayRpcConf)),
	}
}
