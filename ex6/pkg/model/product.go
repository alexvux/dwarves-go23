package model

type Product struct {
	ID       int     `json:"id" binding:"required,min=1"`
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=1"`
	Quantity int     `json:"quantity" binding:"required,min=1"`
}
