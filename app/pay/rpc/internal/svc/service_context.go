package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-shop/app/pay/rpc/internal/config"
	"zero-shop/app/pay/rpc/model"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.ThirdPayModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewThirdPayModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
