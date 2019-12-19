package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-micro-doc/common/constant"
	"go-micro-doc/common/util"
	"io"
	"reflect"
)

const (
	BindFormKey     = "_controllers/bind_form_key"
	ResponseDataKey = "_controllers/response_data_key"
)

func Bind(val interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		panic(`Bind struct can not be a pointer. Example:
	Use: gin.Bind(Struct{}) instead of gin.Bind(&Struct{})
`)
	}

	return func(c *gin.Context) {
		obj := reflect.New(value.Type()).Interface()
		// shouldBindBodyWith 在context存储了 BodyBytesKey, 需要在后置的middleware当中使用该值
		if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil && err != io.EOF {
			util.Fail(c, constant.ErrorParams)
			c.Abort()
			return
		}
		c.Set(BindFormKey, obj)
		c.Next()
	}
}
