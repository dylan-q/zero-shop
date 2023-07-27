package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/app/goods/api/internal/config"
	"zero-shop/app/goods/api/internal/middleware"
	"zero-shop/app/goods/rpc/goods"
)

type ServiceContext struct {
	Config        config.Config
	GoodsRpc      goods.Goods
	LogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		GoodsRpc:      goods.NewGoods(zrpc.MustNewClient(c.GoodsRpcConf)),
		LogMiddleware: middleware.NewLogMiddleware().Handle,
	}
}
