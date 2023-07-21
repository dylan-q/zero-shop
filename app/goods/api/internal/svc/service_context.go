package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/app/goods/api/internal/config"
	"zero-shop/app/goods/rpc/goods"
)

type ServiceContext struct {
	Config   config.Config
	GoodsRpc goods.Goods
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsRpc: goods.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
	}
}
