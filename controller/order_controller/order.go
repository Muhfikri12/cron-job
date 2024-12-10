package ordercontroller

import (
	"cron-job/helper"
	"cron-job/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController interface {
	GetListOrder(c *gin.Context)
}

type orderController struct {
	service *service.AllService
	log     *zap.Logger
}

func NewOrderController(service *service.AllService, log *zap.Logger) OrderController {
	return &orderController{service, log}
}

func (oc *orderController) GetListOrder(c *gin.Context) {

	orders, err := oc.service.Order.GetListOrder()
	if err != nil {
		helper.Responses(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "Successfully Retrieved Orders", orders)
}
