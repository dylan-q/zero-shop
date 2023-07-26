package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/app/order/rpc/order"
	"zero-shop/app/pay/api/internal/config"
	"zero-shop/app/pay/rpc/pay"
)

type ServiceContext struct {
	Config   config.Config
	PayRpc   pay.Pay
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		PayRpc:   pay.NewPay(zrpc.MustNewClient(c.PayRpcConf)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
