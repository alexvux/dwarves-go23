package repo

import (
	"fmt"

	"github.com/alexvux/dwarves-go23/ex6/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
)

func AddItemToCart(item model.Item) error {
	// check if product exists
	if _, err := findProductByID(item.ID); err != nil {
		return err
	}
	// check if item already exists in cart, then update quantity
	if idx, err := findItemByID(item.ID); err == nil {
		DB.Cart[idx].Quantity += item.Quantity
		return nil
	}
	// else add new item to cart
	DB.Cart = append(DB.Cart, item)
	return nil
}

func DeleteItemFromCart(item model.Item) error {
	idx, err := findItemByID(item.ID)
	if err != nil {
		return err
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
	pidToCheckout := []int{}

	for _, item := range DB.Cart {
		// check if product exists
		pid, err := findProductByID(item.ID)
		if err != nil {
			return "", wrapErrorWithItemID(item.ID, err)
		}
		// check if product is in stock
		product := DB.Products[pid]
		if product.Quantity < item.Quantity {
			return "", wrapErrorWithItemID(item.ID, constant.ErrOutOfStock)
		}
		// update reciept and save product to checkout
		totalPrice += product.Price * float64(item.Quantity)
		reciept += fmt.Sprintf("\t%s x %d = %.2f\n", product.Name, item.Quantity, product.Price*float64(item.Quantity))
		pidToCheckout = append(pidToCheckout, pid)
	}
	// update cart and products
	for _, idx := range pidToCheckout {
		DB.Products[idx].Quantity -= DB.Cart[idx].Quantity
	}
	DB.Cart = []model.Item{}
	return fmt.Sprintf("%s\tTotal: %.2f", reciept, totalPrice), nil
}

func findItemByID(id int) (int, error) {
	for idx, item := range DB.Cart {
		if id == item.ID {
			return idx, nil
		}
	}
	return 0, constant.ErrItemNotFound
}

func wrapErrorWithItemID(id int, err error) error {
	return fmt.Errorf("item with id %d: %w", id, err)
}
