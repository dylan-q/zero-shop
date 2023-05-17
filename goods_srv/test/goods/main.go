package main

import (
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"
	"strings"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

//	func TestBatchGetGoods() {
//		rsp, err := brandClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
//			Id: []int32{421, 422, 423},
//		})
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(rsp.Total)
//		for _, good := range rsp.Data {
//			fmt.Println(good.Name, good.ShopPrice)
//		}
//	}
//
//	func TestGetGoodsDetail() {
//		rsp, err := brandClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
//			Id: 421,
//		})
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(rsp.Name)
//		fmt.Println(rsp.DescImages)
//	}
//
//	func Init() {
//		var err error
//		conn, err = grpc.Dial("192.168.0.112:14566", grpc.WithInsecure())
//		if err != nil {
//			panic(err)
//		}
//		brandClient = proto.NewGoodsClient(conn)
//	}
//
//	func TestGetGoodsList() {
//		rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
//			TopCategory: 136688,
//			//PriceMin:    90,
//			//KeyWords: "深海速冻",
//		})
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(rsp.Data)
//		for _, good := range rsp.Data {
//			fmt.Println(good.Name, good.ShopPrice)
//		}
//	}
func InitEs() {
	//初始化连接
	host := fmt.Sprintf("http://%s:%s", "192.168.0.112", "9200")
	var err error
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
}
func main() {
	InitEs()
	//Init()
	//TestCreateUser()
	//TestGetGoodsList()
	//TestBatchGetGoods()
	//TestGetGoodsDetail()

	//conn.Close()
}
