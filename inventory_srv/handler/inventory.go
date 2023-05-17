package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"mxshop_srvs/inventory_srv/global"
	"mxshop_srvs/inventory_srv/model"
	"mxshop_srvs/inventory_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (i *InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存， 如果我要更新库存
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	inv.Goods = req.GoodsId
	inv.Stocks = req.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (i *InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}
func (i *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {

	client := goredislib.NewClient(&goredislib.Options{
		Addr: "127.0.0.1:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
	rs := redsync.New(pool)
	tx := global.DB.Begin()
	sellDetail := model.StockSellDetail{
		OrderSn: req.OrderSn,
		Status:  1,
		Detail:  nil,
	}
	var detail []model.GoodsDetail
	for _, goodsInfo := range req.GoodsInfo {
		detail = append(detail, model.GoodsDetail{
			Goods: goodsInfo.GoodsId,
			Num:   goodsInfo.Num,
		})
		var inv model.Inventory
		mutex := rs.NewMutex(fmt.Sprintf("goods_%d", goodsInfo.GoodsId))
		if err := mutex.Lock(); err != nil {
			return nil, status.Errorf(codes.Internal, "获取分布式锁失败"+err.Error())
		}
		if result := global.DB.Where("goods = ?", goodsInfo.GoodsId).First(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}
		if inv.Stocks < goodsInfo.Num {
			tx.Rollback()
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		inv.Stocks -= goodsInfo.Num
		if result := tx.Model(&model.Inventory{}).Select("stocks", "version").Where("goods = ? and version = ?", goodsInfo.GoodsId, inv.Version).Updates(model.Inventory{
			Stocks: inv.Stocks, Version: inv.Version + 1,
		}); result.RowsAffected == 0 {
			zap.S().Info("库存扣减失败")
		}
		tx.Save(&inv)
		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Errorf(codes.Internal, "释放分布式锁失败")
		}
	}
	sellDetail.Detail = detail
	if result := tx.Create(&sellDetail); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "保存库存扣减失败")
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}
func (i *InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//库存归还： 1：订单超时归还 2. 订单创建失败，归还之前扣减的库存 3. 手动归还
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() //回滚之前的操作
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}

		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
		inv.Stocks += goodInfo.Num
		tx.Save(&inv)
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}

func AutoReback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	type OrderInfo struct {
		OrderSn string
	}
	for i := range msgs {
		var orderInfo OrderInfo
		err := json.Unmarshal(msgs[i].Body, &orderInfo)
		if err != nil {
			zap.S().Errorf("解析json失败：%v \n", msgs[i].Body)
			return consumer.ConsumeSuccess, nil
		}
		tx := global.DB.Begin()
		var sellDetail model.StockSellDetail
		if result := tx.Where(&model.StockSellDetail{
			OrderSn: orderInfo.OrderSn, Status: 1,
		}).First(&sellDetail); result.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}
		for _, orderGoods := range sellDetail.Detail {
			if result := tx.Model(&model.Inventory{}).Where(&model.Inventory{Goods: orderGoods.Goods}).Update("stocks", gorm.Expr("stocks + ?", orderGoods.Num)); result.RowsAffected == 0 {
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil
			}
		}
		sellDetail.Status = 2
		if res := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{OrderSn: orderInfo.OrderSn}).Update("status", 2); res.RowsAffected == 0 {
			tx.Rollback()
			return consumer.ConsumeRetryLater, nil
		}
		tx.Commit()
		return consumer.ConsumeSuccess, nil
	}

	return consumer.ConsumeSuccess, nil
}
