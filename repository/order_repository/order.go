package orderrepository

import (
	"cron-job/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepo interface {
	GetListOrder() (*[]model.Checkout, error)
}

type orderRepo struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewOrderRepo(db *gorm.DB, log *zap.Logger) OrderRepo {
	return &orderRepo{DB: db, Log: log}
}

func (o *orderRepo) GetListOrder() (*[]model.Checkout, error) {
	orders := []model.Checkout{}

	result := o.DB.Find(&orders)

	if result.Error != nil {
		o.Log.Error("Failed to fetch orders", zap.Error(result.Error))
		return nil, result.Error
	}

	if len(orders) == 0 {
		o.Log.Info("No orders found")
		return nil, fmt.Errorf("no history order")
	}

	o.Log.Info("Successfully fetched orders", zap.Int("count", len(orders)))
	return &orders, nil
}
