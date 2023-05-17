package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderListener struct {
	ID int32
}

func (o *OrderListener) ExecuteLocalTransaction(addr *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second * 3)
	o.ID = 10
	fmt.Println("执行本地逻辑成功")
	return primitive.CommitMessageState
}
func (o *OrderListener) CheckLocalTransaction(addr *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq 消息回查")
	time.Sleep(time.Second * 3)
	return primitive.CommitMessageState
}

// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
// method will be invoked to get local transaction status.
//CheckLocalTransaction(*MessageExt) LocalTransactionState

func main() {
	order := OrderListener{}
	p, err := rocketmq.NewTransactionProducer(
		&order,
		producer.WithNameServer([]string{"192.168.0.112:9876"}),
	)
	if err != nil {
		panic(err)
	}
	if err := p.Start(); err != nil {
		panic(err)
	}
	transaction, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("transaction", []byte("这是失败的事务消息")))
	if err != nil {
		return
	}
	if err != nil {
		fmt.Printf("发生失败-------%s\n", err.Error())
	} else {
		fmt.Printf("发生成功-------%s\n", transaction.String())
	}

	fmt.Println("主函数==", order.ID)
	time.Sleep(time.Hour)
	if err := p.Shutdown(); err != nil {
		panic("关闭失败")
	}
}
