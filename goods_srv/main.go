package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/handler"
	"mxshop_srvs/goods_srv/initialize"
	"mxshop_srvs/goods_srv/proto"
	"mxshop_srvs/goods_srv/utils"
	"mxshop_srvs/goods_srv/utils/register/consul"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ip := flag.String("ip", "", "ip地址")
	port := flag.Int("port", 0, "端口号")
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitEs()

	flag.Parse()

	if *port == 0 {
		*port, _ = utils.GetFreePort()
	}
	if *ip == "" {
		*ip = utils.GetIP()
	}
	zap.S().Info("IP: ", *ip)
	zap.S().Info("PORT: ", *port)
	server := grpc.NewServer()
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		panic(err)
	}

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = registerClient.Register(*ip, *port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
