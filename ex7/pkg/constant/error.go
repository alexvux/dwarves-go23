package constant

import "errors"

var (
	ErrEmptyProductList    = errors.New("product list is empty")
	ErrProductAlreadyExist = errors.New("product already exist")
	ErrProductNotFound     = errors.New("product not found")
	ErrItemNotFound        = errors.New("item not found")
	ErrEmptyCart           = errors.New("cart is empty")
	ErrOutOfStock          = errors.New("out of stock")
)
