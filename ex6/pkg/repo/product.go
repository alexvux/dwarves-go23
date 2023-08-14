package repo

import (
	"github.com/alexvux/dwarves-go23/ex6/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
)

func GetAllProducts() ([]model.Product, error) {
	if len(DB.Products) == 0 {
		return []model.Product{}, constant.ErrEmptyProductList
	}
	return DB.Products, nil
}

func AddProduct(product model.Product) error {
	if findProductIdxByID(product.ID) == -1 {
		DB.Products = append(DB.Products, product)
		return nil
	}
	return constant.ErrProductAlreadyExist
}

func UpdateProduct(id int, product model.Product) error {
	idx := findProductIdxByID(id)
	if idx == -1 {
		return constant.ErrProductNotFound
	}
	DB.Products[idx] = product
	return nil
}

func DeleteProduct(id int) error {
	idx := findProductIdxByID(id)
	if idx == -1 {
		return constant.ErrProductNotFound
	}
	last := len(DB.Products) - 1
	DB.Products[idx] = DB.Products[last]
	DB.Products = DB.Products[:last]
	return nil
}

func findProductIdxByID(id int) int {
	for i, p := range DB.Products {
		if p.ID == id {
			return i
		}
	}
	return -1
}
