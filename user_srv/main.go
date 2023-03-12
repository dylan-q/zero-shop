package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/proto"
	"net"
)

func main() {
	ip := flag.String("ip", "0.0.0.0", "ip地址")
	port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	fmt.Println("ip :", *ip)
	fmt.Println("port :", *port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		panic(err)
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
