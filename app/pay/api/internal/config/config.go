package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	PayRpcConf   zrpc.RpcClientConf
	OrderRpcConf zrpc.RpcClientConf
	MicroService struct {
		Driver   string
		Target   string
		EndPoint string
	}
}
