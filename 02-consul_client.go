package main

import (
	"context"
	"day01/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	//初始话consul配置
	consulConfig := api.DefaultConfig()
	//创建consul对象
	consulClient, _ := api.NewClient(consulConfig)
	//服务发现，从consul上获取健康状态的服务
	services, _, _ := consulClient.Health().Service("grpc And Consul", "grpc", true, nil)

	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	//1.连接服务
	grpcConn, _ := grpc.Dial(addr, grpc.WithInsecure())
	//2.初始化 grpc客户端
	grpcClient := pb.NewHelloClient(grpcConn)
	//3.调用远程函数
	person := pb.Person{
		Name: "jiji!!!",
		Age:  20,
	}
	p, _ := grpcClient.SayHello(context.TODO(), &person)
	fmt.Println("your info----->", p)
}
