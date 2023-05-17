package global

import (
	"github.com/apache/rocketmq-client-go/v2"
	"gorm.io/gorm"
	"mxshop_srvs/order_srv/config"
	"mxshop_srvs/order_srv/proto"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
	RocketMq           rocketmq.TransactionProducer
)
