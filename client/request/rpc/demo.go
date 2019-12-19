package rpc

import (
	"context"
	go_micro_demo "go-micro-doc/proto/demo"
	"sync"
)

type demo struct {
	service go_micro_demo.DemoService
}

var demoRpc *demo
var demoOnce sync.Once

/**
 * 获取实例
 */
func NewDemo() *demo {
	demoOnce.Do(func() {
		demoRpc = &demo{service: go_micro_demo.NewDemoService("", NewMicroClient("", ""))}
	})
	return demoRpc
}

/**
 *
 */
func (d *demo) Hello(name string) (*go_micro_demo.HelloResponse, error) {
	resp, err := d.service.Hello(context.TODO(), &go_micro_demo.NameRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

