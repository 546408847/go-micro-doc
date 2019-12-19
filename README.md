## go-micro-doc
go-micro 微服务部署安装使用学习教程


## 搭建开发环境

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
startmicroweb='micro --registry=etcd --registry_address=127.0.0.1:8500 web' //micro网关的web可视化界面，默认监听8082端口，访问http://localhost:8082查看服务
createproto='protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=.'
```

## 开发指南

### protobuf - go

1. [官方手册](https://developers.google.com/protocol-buffers/docs/gotutorial)
2. [Protobuf 终极教程](https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/)
3. 生成 proto
   ```
   cd /proto/demo
   使用设置好的alias
   createproto test.proto
   ```

### micro 框架
1. [官方手册](https://micro.mu/docs/cn/index.html)
2. [编写服务](https://micro.mu/docs/cn/writing-a-go-service.html)
3. [官方用例](https://github.com/micro/examples)

### 基础学习
1. [ web入口使用gin](https://github.com/micro/examples/blob/master/greeter/api/gin/gin.go)

### 目录结构
```
│───client(web项目入口)
|   |───handler(外部请求接口)
|   |───middleware(中间件)    
|   |───main.go(项目启动文件)
│───srv(rpc入口) 
|   |───handler(微服务接口)   
|   |───main.go(项目启动文件)   
│───common(共用的工具类) 
│───config(配置目录) 
│───proto(proto文件)   

在真实的开发情况下，往往client，srv，proto这3个目录会分别建立三个项目，由于是demo，所以都放一个项目里面了
```

### 启动项目
1. 启动 etcd
   ```
   brew services start etcd
   ```
2. 启动 micro
   ```
   使用先前设置好的alias
   startmicroapi
   startmicroweb
   ```
3. 启动 srv
   ```
   go run srv/main.go
   ```
4. 启动 client
   ```
   go run client/main.go
   ```
5. 发起请求
   ```
   curl -X POST \
     http://localhost:8080/demo/demo/hello \
     -H 'Accept: */*' \
     -H 'Accept-Encoding: gzip, deflate' \
     -H 'Cache-Control: no-cache' \
     -H 'Connection: keep-alive' \
     -H 'Content-Length: 20' \
     -H 'Content-Type: application/json' \
     -H 'Host: localhost:8080' \
     -H 'Postman-Token: 8ff4821a-b3b8-4d3b-be6b-0ba7a9a8723d,35f6c4bb-0ddc-4383-86c4-846a0bd825bd' \
     -H 'User-Agent: PostmanRuntime/7.19.0' \
     -H 'cache-control: no-cache' \
     -d '{
   	"name":"go"
   }'
   ```
