package main

import (
	"fmt"
	"go-micro-doc/common/component"
	"go-micro-doc/common/constant"
	"go-micro-doc/common/model"
	"go-micro-doc/common/util"
	"go-micro-doc/config"
	"go-micro-doc/srv/handler"
)

func init() {
	//初始化配置文件
	config.InitConfig("service", util.GetArg(constant.ArgEnv, "dev"), util.GetArg(constant.ArgConfig, ""))
}

func main() {
	//注册数据库连接
	model.RegisterDB()
	defer model.CloseDB()

	//注册服务
	service := component.NewSvrService(constant.RegistryEtcd, constant.SysSvrServiceName)
	//初始化插件、命令行参数
	service.Init()
	//绑定处理器
	handler.RegisterHandler(service)
	//启动
	if err := service.Run(); err != nil {
		panic(fmt.Errorf("fatal error: run service: %s\n", err))
	}
}
