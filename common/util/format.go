package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/errors"
	"go-micro-doc/common/constant"
	"net/http"
)

/**
 * 成功
 */
func Success(ctx *gin.Context, data interface{}) {
	d := gin.H{
		"code":    constant.ErrorOk,
		"message": constant.ErrorMessage(constant.ErrorOk),
		"data":    data,
	}
	ctx.JSON(http.StatusOK, d)
}

/**
 * 失败
 */
func Fail(ctx *gin.Context, errs ...interface{}) {
	code := constant.ErrorSystemError
	message := constant.ErrorMessage(code)
	for _, e := range errs {
		switch e.(type) {
		case int:
			code = e.(int)
			message = constant.ErrorMessage(code)
		case string:
			message = e.(string)
		case *errors.Error:
			err := e.(*errors.Error)
			code = int(err.Code)
			message = err.Status
		case error:
			message = e.(error).Error()
		}
	}
	data := gin.H{
		"code":    code,
		"message": message,
		"data":    gin.H{},
	}

	ctx.JSON(http.StatusOK, data)
}

/**
 * 错误
 * 必须:错误编号
 * 可选:系统错误，自定义业务错误信息
 */
func Error(code int, errs ...interface{}) *errors.Error {
	err := &errors.Error{
		Code:   int32(code),                 //业务错误编号
		Status: constant.ErrorMessage(code), //业务错误信息
	}
	for _, e := range errs {
		switch e.(type) {
		case string:
			err.Status = e.(string)
		case *errors.Error:
			*err = *e.(*errors.Error)
		case error:
			detail := e.(error).Error() //系统错误信息
			if ue := json.Unmarshal([]byte(detail), err); ue != nil {
				err.Detail = detail
			}
		}
	}

	return err
}
