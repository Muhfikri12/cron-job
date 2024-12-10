package orderservice

import (
	"cron-job/model"
	"cron-job/repository"
	orderrepository "cron-job/repository/order_repository"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderService interface {
	GetListOrder() (*[]model.Checkout, error)
}

type orderService struct {
	Repo *repository.AllRepository
	Log  *zap.Logger
}

func NewOrderService(Repo *repository.AllRepository, Log *zap.Logger) OrderService {
	return &orderService{Repo, Log}
}

func (os *orderService) GetListOrder() (*[]model.Checkout, error) {

	orders, err := os.Repo.Order.GetListOrder()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func StartCronJob(db *gorm.DB, logger *zap.Logger) {
	c := cron.New()

	_, err := c.AddFunc("* * * * *", func() {
		repo := orderrepository.NewOrderRepo(db, logger)
		orders, err := repo.GetListOrder()
		if err != nil {
			logger.Error("Failed to fetch orders", zap.Error(err))
			return
		}

		if len(*orders) == 0 {
			logger.Warn("No orders found")
			return
		}

		logger.Info("Fetched orders", zap.Int("orderCount", len(*orders)))

		file := excelize.NewFile()
		sheetName := "Orders"
		i, _ := file.NewSheet(sheetName)

		headers := []string{"ID", "UserID", "TotalAmount", "Payment", "ShippingPayment", "CreatedAt", "UpdatedAt"}
		for i, header := range headers {
			cell := fmt.Sprintf("%s1", string(rune('A'+i)))
			file.SetCellValue(sheetName, cell, header)
		}

		for rowIndex, order := range *orders {
			values := []interface{}{order.ID, order.UserID, order.TotalAmount, order.Payment, order.ShippingPayment, order.CreatedAt, order.UpdatedAt}
			for colIndex, value := range values {
				cell := fmt.Sprintf("%s%d", string(rune('A'+colIndex)), rowIndex+2)
				file.SetCellValue(sheetName, cell, value)
			}
		}

		file.SetActiveSheet(i)

		fileName := fmt.Sprintf("orders_%s.xlsx", time.Now().Format("20060102_150405"))
		if err := file.SaveAs(fileName); err != nil {
			logger.Error("Failed to save Excel file", zap.Error(err))
			return
		}

		logger.Info("Excel file created successfully", zap.String("fileName", fileName))
	})

	if err != nil {
		logger.Error("Failed to schedule cron job", zap.Error(err))
		return
	}

	c.Start()
	logger.Info("Cron job started successfully, running every 1 minute")

	select {}
}
