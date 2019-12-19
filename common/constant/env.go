package constant

import "github.com/gin-gonic/gin"

const (
	EnvDev  = "dev"
	EnvTest = "test"
	EnvProd = "prod"
)

/**
 * 获取gin运行模式
 */
func GetGinMode(env string) string {
	switch env {
	case EnvTest:
		return gin.TestMode
	case EnvProd:
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}
