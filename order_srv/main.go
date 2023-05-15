package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop_srvs/goods_srv/handler"
	"mxshop_srvs/goods_srv/initialize"
	"mxshop_srvs/goods_srv/proto"
)

func main() {

	ip := flag.String("ip", "", "ip地址")
	port := flag.Int("port", 0, "端口号")
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

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

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", *ip, *port), //监听的服务地址
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serviceId
	registration.Port = *port
	registration.Tags = global.ServerConfig.Tags
	registration.Address = *ip
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	//client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
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
	if err = client.Agent().ServiceDeregister(serviceId); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
