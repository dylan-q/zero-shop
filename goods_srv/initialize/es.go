package initialize

import (
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"strings"
)

func InitEs() {
	//初始化连接
	host := fmt.Sprintf("http://%s:%d", global.ServerConfig.EsInfo.Host, global.ServerConfig.EsInfo.Port)
	//logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
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
