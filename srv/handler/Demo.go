package handler

import (
	go_micro_demo "go-micro-doc/proto/demo"
	"context"
)

type Demo struct {

}

func (d *Demo) Hello(ctx context.Context, req *go_micro_demo.NameRequest, resp *go_micro_demo.HelloResponse) error{
	//todo db curd
	resp.Content = req.Name + "欢迎你，成功调用rpc"
	return nil
}