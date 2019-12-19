package handler

import (
	"github.com/gin-gonic/gin"
	"go-micro-doc/client/middleware"
	"go-micro-doc/client/middleware/form"
	"go-micro-doc/common/constant"
)

func RegisterRoute(router *gin.Engine) {
	e := router.Group(constant.SysRoutePre)
	{
		e.POST("/demo/hello", middleware.Bind(form.Demo{}), new(Demo).Hello)
	}
}
