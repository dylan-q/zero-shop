package initialize

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
	"mxshop_srvs/order_srv/global"
	"mxshop_srvs/order_srv/handler"
)

func InitMq() {
	orderListen := handler.OrderListener{}
	p, err := rocketmq.NewTransactionProducer(
		&orderListen,
		producer.WithNameServer([]string{"192.168.0.112:9876"}),
		producer.WithGroupName("shop-inventory"),
	)
	if err != nil {
		zap.S().Errorf("生成producer失败-------%s\n", err.Error())
		panic(err)
	}
	if err := p.Start(); err != nil {
		zap.S().Errorf("启动producer失败-------%s\n", err.Error())
		panic(err)
	}
	global.RocketMq = p
}
