package handler

import (
	"context"
	goMicroDemo "go-micro-doc/proto/demo"
)

type Demo struct {

}

func (d *Demo) Hello(ctx context.Context, req *goMicroDemo.NameRequest, resp *goMicroDemo.HelloResponse) error{
	//todo db curd
	resp.Content = req.Name + "!欢迎你，成功调用rpc"
	return nil
}