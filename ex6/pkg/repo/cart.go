package repo

import (
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
		LocalDB.Cart[idx].Quantity += item.Quantity
		return nil
	}
	// else add new item to cart
	LocalDB.Cart = append(LocalDB.Cart, item)
	return nil
}

func DeleteItemFromCart(item model.Item) error {
	idx, err := findItemByID(item.ID)
	if err != nil {
		return err
	}
	last := len(LocalDB.Cart) - 1
	LocalDB.Cart[idx] = LocalDB.Cart[last]
	LocalDB.Cart = LocalDB.Cart[:last]
	return nil
}

func findItemByID(id int) (int, error) {
	for i, p := range LocalDB.Cart {
		if id == p.ID {
			return i, nil
		}
	}
	return 0, constant.ErrItemNotFound
}
