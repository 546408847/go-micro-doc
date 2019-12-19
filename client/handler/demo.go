package handler

import (
	"github.com/gin-gonic/gin"
	"go-micro-doc/client/middleware"
	"go-micro-doc/client/middleware/form"
	"go-micro-doc/client/request/rpc"
	"go-micro-doc/common/constant"
	"go-micro-doc/common/util"
)

type Demo struct {

}

func (d *Demo) Hello(c *gin.Context) {
	param := c.MustGet(middleware.BindFormKey).(*form.Demo)
	//获取订单基本信息
	message, err := rpc.NewDemo().Hello(param.Name)
	if err != nil {
		util.Fail(c, constant.ErrorRpcRequest)
		return
	}
	util.Success(c, message)
	return
}