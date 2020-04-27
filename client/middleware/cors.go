package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerKeys []string
		for k, v := range c.Request.Header {
			if k == "Access-Control-Request-Headers" {
				for _, v1 := range v {
					headerKeys = append(headerKeys, v1)
				}
			}
		}
		allowHeaders := []string{"*"}
		if len(headerKeys) != 0 {
			allowHeaders = headerKeys
		}

		//如果需要跨域携带cookie，origin需要指定值，不能使用"*"
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}

		cors.New(cors.Config{
			AllowOrigins:     []string{origin},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEADER", "OPTIONS"},
			AllowHeaders:     allowHeaders,
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})(c)
	}
}
