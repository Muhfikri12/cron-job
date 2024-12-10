package router

import (
	"cron-job/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.IntegrationContext) *gin.Engine {

	r := gin.Default()

	r.POST("/login", ctx.Ctl.Auth.Login)
	r.GET("/order", ctx.Ctl.Order.GetListOrder)

	return r
}
