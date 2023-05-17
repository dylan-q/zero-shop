package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/inventory_srv/proto"
	"sync"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("192.168.0.112:10762", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}
func TestSetInv(goodsId, Num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}
func TestSell(wg *sync.WaitGroup) {
	/*
		1. 第一件扣减成功： 第二件： 1. 没有库存信息 2. 库存不足
		2. 两件都扣减成功
	*/
	defer wg.Done()
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},
			//{GoodsId: 422, Num: 30},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("库存扣减成功")
}
func main() {
	var wg sync.WaitGroup
	Init()
	//TestCreateUser()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go TestSell(&wg)
	}
	wg.Wait()
	//TestBatchGetGoods()
	//TestGetGoodsDetail()

	conn.Close()
}
