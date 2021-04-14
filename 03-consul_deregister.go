package main

import "github.com/hashicorp/consul/api"

func main() {
	//1.初始话consul配置
	consulConfig := api.DefaultConfig()
	//2.创建consul对象
	consulClient, _ := api.NewClient(consulConfig)
	//3.注销服务
	consulClient.Agent().ServiceDeregister("bj38")
}
