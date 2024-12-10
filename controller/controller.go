package controller

import (
	authcontroller "cron-job/controller/auth_controller"
	ordercontroller "cron-job/controller/order_controller"
	"cron-job/database"
	"cron-job/service"

	"go.uber.org/zap"
)

type AllController struct {
	Auth  authcontroller.AuthHadler
	Order ordercontroller.OrderController
}

func NewAllController(service *service.AllService, log *zap.Logger, cfg *database.Cache) AllController {
	return AllController{
		Auth:  authcontroller.NewUserHandler(service, log, cfg),
		Order: ordercontroller.NewOrderController(service, log),
	}
}
