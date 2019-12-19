package rpc

import (
	"go-micro-doc/common/constant"
	"go-micro-doc/config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/registry/consul"
	"strings"
	"time"
)

/**
 * 获取micro rpc请求client实例
 */
func NewMicroClient(client, key string, opts ...micro.Option) client.Client {
	registryConfig := config.AppConfig.Registry
	opt := registry.Option(func(opts *registry.Options) {
		address := registryConfig.Address
		if key != "" {
			address = registryConfig.Client[key]
		}
		opts.Addrs = strings.Split(address, ",")
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
		micro.Version(registryConfig.Version),                                        //微服务版本
		micro.Registry(registryOpt),                                                  //注册微服务
		micro.RegisterTTL(time.Second * time.Duration(registryConfig.Ttl)),           //微服务发现组件中的节点存活时间
		micro.RegisterInterval(time.Second * time.Duration(registryConfig.Interval)), //微服务发现组件中的节点多久刷新自己，来避免自己被 register_ttl 掉
		micro.Metadata(registryConfig.MetaData),
	}

	//合并选项
	opts = append(defaultOpts, opts...)

	//获取微服务实例
	service := micro.NewService(opts...)

	return service.Client()
}
