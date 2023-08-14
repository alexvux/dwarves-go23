package repo

import (
	"fmt"

	"github.com/alexvux/dwarves-go23/ex6/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
)

func AddItemToCart(item model.Item) error {
	// check if product exists
	if findProductIdxByID(item.ID) == -1 {
		return constant.ErrProductNotFound
	}
	// check if item already exists in cart, then update quantity
	idx := findItemIdxByID(item.ID)
	if idx != -1 {
		DB.Cart[idx].Quantity += item.Quantity
		return nil
	}
	// else add new item to cart
	DB.Cart = append(DB.Cart, item)
	return nil
}

func DeleteItemFromCart(id int) error {
	idx := findItemIdxByID(id)
	if idx == -1 {
		return constant.ErrItemNotFound
	}
	last := len(DB.Cart) - 1
	DB.Cart[idx] = DB.Cart[last]
	DB.Cart = DB.Cart[:last]
	return nil
}

func Checkout() (string, error) {
	if len(DB.Cart) == 0 {
		return "", constant.ErrEmptyCart
	}

	totalPrice := float64(0)
	reciept := "Reciept:\n"
	idxToCheckout := []int{}
	productToCheckout := []model.Product{}

	for _, item := range DB.Cart {
		// check if product exists
		idx := findProductIdxByID(item.ID)
		if idx == -1 {
			return "", wrapErrorWithItemID(item.ID, constant.ErrProductNotFound)
		}
		// check if product is in stock
		product := DB.Products[idx]
		if product.Quantity < item.Quantity {
			return "", wrapErrorWithItemID(item.ID, constant.ErrOutOfStock)
		}
		// update reciept and save product to checkout
		totalPrice += product.Price * float64(item.Quantity)
		reciept += fmt.Sprintf("\t%s x %d = %.2f\n", product.Name, item.Quantity, product.Price*float64(item.Quantity))
		idxToCheckout = append(idxToCheckout, idx)
		product.Quantity -= item.Quantity
		productToCheckout = append(productToCheckout, product)
	}
	// update cart and products // dang sai
	for i, idx := range idxToCheckout {
		DB.Products[idx] = productToCheckout[i]
	}
	DB.Cart = []model.Item{}
	return fmt.Sprintf("%sTotal: %.2f", reciept, totalPrice), nil
}

func findItemIdxByID(id int) int {
	for i, item := range DB.Cart {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func wrapErrorWithItemID(id int, err error) error {
	return fmt.Errorf("item with id %d: %w", id, err)
}
