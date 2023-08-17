package model

import (
	"time"

	"github.com/alexvux/dwarves-go23/ex7/pkg/constant"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status       constant.OrderStatus `json:"status" binding:"required"`
	CheckoutDate time.Time            `json:"checkout_date"`
	Items        []OrderItem          `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id" binding:"required,min=1"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}
