package routers

import (
	"go-generator/logger"
	"go-generator/settings"

	"github.com/gin-gonic/gin"
)

func Setup(cfg *settings.AppConfig) *gin.Engine {
	gin.SetMode(cfg.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecover())
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  "Page Not Found.If you have any questions,please contact developer#xxcheng.cn~",
		})
	})
	return r
}
