package model

type Item struct {
	ID       int `json:"id" binding:"required,min=1"`
	Quantity int `json:"quantity" binding:"required,min=1"`
}
