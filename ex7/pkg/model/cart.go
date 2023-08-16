package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending OrderStatus = "pending"
	OrderStatusPaid    OrderStatus = "paid"
)

type Order struct {
	gorm.Model
	OrderStatus  string      `json:"order_status" binding:"required"`
	CheckoutDate time.Time   `json:"checkout_date"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   int `json:"order_id" binding:"required,min=1"`
	ProductID int `json:"product_id" binding:"required,min=1"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}
