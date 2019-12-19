package component

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro-doc/common/constant"
	"go-micro-doc/config"
	"strings"
	"time"
	microTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

/**
 * 获取micro web微服务实例
 */
func NewWebService(client, serviceName string, opts ...web.Option) web.Service {
	registryConfig := config.AppConfig.Registry
	opt := registry.Option(func(opts *registry.Options) {
		opts.Addrs = strings.Split(registryConfig.Address, ",")
	})
	var registryOpt registry.Registry
	switch client {
	case constant.RegistryConsul:
		registryOpt = consul.NewRegistry(opt)
	default:
		registryOpt = etcd.NewRegistry(opt)
	}

	//默认选项
	defaultOpts := []web.Option{
		web.Name(serviceName),                                                      //微服务名称
		web.Version(registryConfig.Version),                                        //微服务版本
		web.Registry(registryOpt),                                                  //注册微服务
		web.RegisterTTL(time.Second * time.Duration(registryConfig.Ttl)),           //微服务发现组件中的节点存活时间
		web.RegisterInterval(time.Second * time.Duration(registryConfig.Interval)), //微服务发现组件中的节点多久刷新自己，来避免自己被 register_ttl 掉
		web.Metadata(registryConfig.MetaData),                                      //微服务自身元数据，会上报给服务发现组件。是一个 key-value 列表
		web.Flags(cli.StringFlag{ //自定义命令行参数
			Name:  "env",
			Value: "dev",
			Usage: "app runtime environment. the default value is dev",
		}),
		web.Flags(cli.StringFlag{ //自定义命令行参数
			Name:  "config",
			Value: "",
			Usage: "app config file path. the default value is get by env",
		}),
		//micro.Address("127.0.0.1:8080"), //微服务服务地址，不需要指定，系统自动分配未使用的端口
	}

	//合并选项
	opts = append(defaultOpts, opts...)

	//获取微服务实例，选项可使用命令行参数调整
	service := web.NewService(opts...)

	return service
}

/**
 * 获取micro微服务实例
 */
func NewSvrService(client, serviceName string, opts ...micro.Option) micro.Service {
	registryConfig := config.AppConfig.Registry
	opt := registry.Option(func(opts *registry.Options) {
		opts.Addrs = strings.Split(registryConfig.Address, ",")
	})
	var registryOpt registry.Registry
	switch client {
	case constant.RegistryConsul:
		registryOpt = consul.NewRegistry(opt)
	default:
		registryOpt = etcd.NewRegistry(opt)
	}

	//默认选项
	defaultOpts := []micro.Option{
		micro.Name(serviceName),                                                      //微服务名称
		micro.Version(registryConfig.Version),                                        //微服务版本
		micro.Registry(registryOpt),                                                  //注册微服务
		micro.RegisterTTL(time.Second * time.Duration(registryConfig.Ttl)),           //微服务发现组件中的节点存活时间
		micro.RegisterInterval(time.Second * time.Duration(registryConfig.Interval)), //微服务发现组件中的节点多久刷新自己，来避免自己被 register_ttl 掉
		micro.Metadata(registryConfig.MetaData),                                      //微服务自身元数据，会上报给服务发现组件。是一个 key-value 列表
		micro.Flags(cli.StringFlag{ //自定义命令行参数
			Name:  "env",
			Value: "dev",
			Usage: "app runtime environment. the default value is dev",
		}),
		micro.Flags(cli.StringFlag{ //自定义命令行参数
			Name:  "config",
			Value: "",
			Usage: "app config file path. the default value is get by env",
		}),
		micro.WrapHandler(microTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	}

	//合并选项
	opts = append(defaultOpts, opts...)

	//获取微服务实例
	service := micro.NewService(opts...)

	return service
}

