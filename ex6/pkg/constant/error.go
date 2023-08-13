package constant

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrItemNotFound    = errors.New("item not found")
)
