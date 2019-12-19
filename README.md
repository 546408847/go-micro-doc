# go-micro-doc
go-micro 微服务部署安装使用学习教程



# 搭建开发环境

安装 etcd

```
brew install etcd
```

安装 protobuf

```
https://github.com/protocolbuffers/protobuf/releases
把bin和include文件的内容分别解压到/usr/local/bin和/usr/local/include中

或者
brew install protobuf
```

安装依赖

```sh
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
go get -u github.com/micro/micro
go get -u github.com/micro/go-micro
```

go env

```sh
GO111MODULE="auto"
GOPATH="/Users/user/Applications/go-project"
GOPRIVATE="*.yourjob.com"  //自定义你公司的私有仓库地址
GOPROXY="https://goproxy.cn"
```

各种alias

```sh
startmicroapi='micro --registry=etcd --registry_address=127.0.0.1:8500 api --handler=http' //以etcd注册，来启动http请求的go-micro网关
startmicroweb='micro --registry=etcd --registry_address=127.0.0.1:8500 web' //micro网关的web可视化界面
createproto='protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=.'
```


# 开发指南

### protobuf - go

1. [官方手册](https://developers.google.com/protocol-buffers/docs/gotutorial)
2. [Protobuf 终极教程](https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/)

### micro 框架
1. [官方手册](https://micro.mu/docs/cn/index.html)
2. [编写服务](https://micro.mu/docs/cn/writing-a-go-service.html)
3. [官方用例](https://github.com/micro/examples)

### 基础学习
1. [ web入口使用gin](https://github.com/micro/examples/blob/master/greeter/api/gin/gin.go)