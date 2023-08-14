package constant

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrItemNotFound    = errors.New("item not found")
	ErrEmptyCart       = errors.New("cart is empty")
	ErrOutOfStock      = errors.New("out of stock")
)
