package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=1"`
	Quantity int     `json:"quantity" binding:"required,min=1"`
}
