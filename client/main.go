package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro-doc/client/handler"
	"go-micro-doc/client/middleware"
	"go-micro-doc/common/component"
	"go-micro-doc/common/constant"
	"go-micro-doc/common/model"
	"go-micro-doc/common/util"
	"go-micro-doc/config"
)

func init() {
	//初始化配置文件
	config.InitConfig("service", util.GetArg(constant.ArgConfig, "dev"), util.GetArg(constant.ArgConfig, ""))
}

func main() {
	//注册数据库连接
	model.RegisterDB()
	defer model.CloseDB()

	//注册服务
	service := component.NewWebService(constant.RegistryEtcd, constant.SysApiServiceName)
	//初始化插件、命令行参数
	if err := service.Init(); err != nil {
		panic(fmt.Errorf("fatal error: init web service: %s\n", err))
	}

	//注册路由
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger(), middleware.Cors())
	handler.RegisterRoute(engine)
	service.Handle("/", engine)

	//启动
	if err := service.Run(); err != nil {
		panic(fmt.Errorf("fatal error: run service: %s\n", err))
	}
}
