package model

import (
	"time"

	"gorm.io/gorm"
)

type Checkout struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	UserID          uint            `gorm:"not null" json:"user_id"`
	TotalAmount     float64         `gorm:"not null" json:"total_amount"`
	Payment         string          `gorm:"size:255" json:"payment"`
	PaymentStatus   string          `gorm:"size:255" json:"payment_status"`
	ShippingPayment float64         `gorm:"not null" json:"shipping_payment"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
