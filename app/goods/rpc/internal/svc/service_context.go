package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-shop/app/goods/rpc/internal/config"
	"zero-shop/app/goods/rpc/model"
)

type ServiceContext struct {
	Config     config.Config
	GoodsModel model.GoodsModel
	CateModel  model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GoodsModel: model.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
