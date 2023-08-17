package constant

import "errors"

var (
	ErrProductIDNotMatch = errors.New("id in path and body must be the same")
	ErrEmptyProductList    = errors.New("product list is empty")
	ErrProductAlreadyExist = errors.New("product already exist")
	ErrProductNotFound     = errors.New("product not found")
	ErrProductOutOfStock   = errors.New("product out of stock")
	ErrItemNotFound        = errors.New("item not found")
	ErrEmptyCart           = errors.New("cart is empty")
)
