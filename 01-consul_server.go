package main

import (
	"context"
	"day01/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

//定义类
type Children struct {
}

//绑定类方法，实现接口
func (this *Children) SayHello(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	person.Name = "hello" + person.Name
	return person, nil
}
func main() {
	//1.初始化consul配置
	consulConfig := api.DefaultConfig()
	//2.创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
	}
	//3.告诉consul，即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID:      "bj38",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc And Consul",
		Address: "127.0.0.1",
		Port:    8800,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//4.注册grpc服务到consul
	consulClient.Agent().ServiceRegister(&reg)
	//1.初始化grpc对象
	grpcServer := grpc.NewServer()
	//2.注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))
	//3.设置监听，指定IP/port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务启动.......")
	//4.启动服务
	_ = grpcServer.Serve(listener)
}
