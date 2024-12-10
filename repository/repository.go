package repository

import (
	authrepository "cron-job/repository/auth_repository"
	orderrepository "cron-job/repository/order_repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AllRepository struct {
	Auth  authrepository.AuthRepoInterface
	Order orderrepository.OrderRepo
}

func NewAllRepo(DB *gorm.DB, Log *zap.Logger) *AllRepository {
	return &AllRepository{
		Auth:  authrepository.NewManagementVoucherRepo(DB, Log),
		Order: orderrepository.NewOrderRepo(DB, Log),
	}
}
