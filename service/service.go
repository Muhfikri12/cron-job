package service

import (
	"cron-job/repository"
	authservice "cron-job/service/auth_service"
	orderservice "cron-job/service/order_service"

	"go.uber.org/zap"
)

type AllService struct {
	Auth  authservice.AuthService
	Order orderservice.OrderService
}

func NewAllService(repo *repository.AllRepository, log *zap.Logger) *AllService {
	return &AllService{
		Auth:  authservice.NewManagementVoucherService(repo, log),
		Order: orderservice.NewOrderService(repo, log),
	}
}
