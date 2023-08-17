package repo

import (
	"fmt"

	"github.com/alexvux/dwarves-go23/ex7/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex7/pkg/model"
)

func AddItemToCart(orderItem model.OrderItem) error {
	// check if product exists
	product, err := getProductByID(orderItem.ProductID)
	if err != nil {
		return constant.ErrProductNotFound
	}
	// check if product is in stock
	if product.Quantity < orderItem.Quantity {
		return constant.ErrProductOutOfStock
	}

	order, err := getCurrentOrder()
	if err != nil { // no order found, create new one
		order, err = createOder()
		if err != nil {
			return err
		}
	}
	// check if item exists in cart then update quantity
	for _, item := range order.Items {
		if item.ProductID == orderItem.ProductID {
			item.Quantity += orderItem.Quantity
			return DB.Save(&item).Error
		}
	}
	// else create new item
	orderItem.OrderID = int(order.ID)
	return DB.Create(&orderItem).Error
}

func DeleteItemFromCart(orderItemID int) error {
	var orderItem model.OrderItem
	err := DB.First(&orderItem, orderItemID).Error
	if err != nil {
		return constant.ErrItemNotFound
	}
	return DB.Delete(&orderItem, orderItemID).Error
}

func Checkout() (string, error) {
	order, err := getCurrentOrder()
	if err != nil { // no pending order found
		return "", err
	}
	if len(order.Items) == 0 { // empty cart
		return "", constant.ErrEmptyCart
	}

	receipt := "Receipt:\n"
	total := float64(0)
	for _, item := range order.Items {
		product, err := getProductByID(item.ProductID)
		if err != nil {
			return "", err
		}
		receipt += fmt.Sprintf("\t%s x %d = %.2f\n", product.Name, item.Quantity, product.Price*float64(item.Quantity))
		total += product.Price * float64(item.Quantity)
	}
	receipt += fmt.Sprintf("Total: %.2f\n", total)
	return receipt, DB.Model(&order).Update("status", constant.OrderStatusPaid).Error
}

// get current order with status pending, if not exists create new one
func getCurrentOrder() (model.Order, error) {
	var order model.Order
	err := DB.Preload("Items").Where("status = ?", constant.OrderStatusPending).First(&order).Error
	return order, err
}

func createOder() (model.Order, error) {
	order := model.Order{
		Status: constant.OrderStatusPending,
	}
	err := DB.Create(&order).Error
	return order, err
}
