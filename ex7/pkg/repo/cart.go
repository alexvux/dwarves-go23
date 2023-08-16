package repo

import (
	_ "fmt"

	_ "github.com/alexvux/dwarves-go23/ex7/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex7/pkg/model"
)

func AddItemToCart(item model.OrderItem) error {
	// check if product exists
	// then check if product is in stock
	// check if oder is pending or paid
	// then check if item already in cart
	// then add item to cart
	return nil
}

func DeleteItemFromCart(id int) error {
	// check if item exists in cart
	// then delete item from cart
	return nil
}

func Checkout() (string, error) {
	// check if cart is empty
	// check if product exists
	// check if product is in stock
	// update reciept and save product to checkout
	// update cart and products // dang sai
	return "", nil
}
