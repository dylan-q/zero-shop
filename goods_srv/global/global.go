package global

import (
	elastic "github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
	"mxshop_srvs/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	EsClient *elastic.Client
)
