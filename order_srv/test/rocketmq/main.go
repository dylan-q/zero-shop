package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"192.168.0.112:9876"}),
	)
	if err != nil {
		panic(err)
	}
	if err := p.Start(); err != nil {
		panic(err)
	}
	result, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "test",
		Body:  []byte("Hello RocketMQ Go Client!"),
	})
	if err != nil {
		fmt.Printf("发生失败-------%s\n", err.Error())
	} else {
		fmt.Printf("发生成功-------%s\n", result.String())
	}
	if err := p.Shutdown(); err != nil {
		panic("关闭失败")
	}
}
