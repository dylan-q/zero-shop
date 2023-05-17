package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ //user model 就是user表不加s
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("http://%s:%s", "192.168.0.112", "9200")
	cfg := elastic.Config{
		Username: "elastic",
		Password: "1PxRLhUmCfftP98pCfLP",
		Addresses: []string{
			host,
		},
	}
	global.EsClient, err = elastic.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	exists, err := global.EsClient.Indices.Exists([]string{"goods"})
	if err != nil {
		panic(err)
	}
	if exists.StatusCode != 200 {
		_, err = global.EsClient.Indices.Create(model.EsGoods{}.GetIndexName(), global.EsClient.Indices.Create.WithBody(strings.NewReader(model.EsGoods{}.GetMapping())))
		if err != nil {
			return
		}
	}
	//db.AutoMigrate(&model.Category{}, model.Brands{}, model.Banner{}, model.GoodsCategoryBrand{}, model.Goods{})
	var goods []model.Goods
	db.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}
		payload, err := json.Marshal(esModel)
		if err != nil {
			panic(err)
		}
		_, err = global.EsClient.Create(model.EsGoods{}.GetIndexName(), strconv.Itoa(int(g.ID)), bytes.NewReader(payload))
		if err != nil {
			return
		}
	}
}
