package handler

import (
	"github.com/micro/go-micro"
	go_micro_demo "go-micro-doc/proto/demo"
)

func RegisterHandler(s micro.Service) {
	_ = go_micro_demo.RegisterDemoHandler(s.Server(), new(Demo))
}
