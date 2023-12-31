package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	GoodsRpcConf zrpc.RpcClientConf
	PayRpcConf   zrpc.RpcClientConf
}
