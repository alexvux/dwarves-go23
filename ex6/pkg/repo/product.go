package repo

import (
	"github.com/alexvux/dwarves-go23/ex6/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
)

func GetAllProducts(prodcut model.Product) []model.Product {
	return DB.Products
}

func AddProduct(product model.Product) error {
	if _, err := findProductByID(product.ID); err != nil {
		return err
	}
	DB.Products = append(DB.Products, product)
	return nil
}

func UpdateProduct(product model.Product) error {
	idx, err := findProductByID(product.ID)
	if err != nil {
		return err
	}
	DB.Products[idx] = product
	return nil
}

func DeleteProduct(product model.Product) error {
	idx, err := findProductByID(product.ID)
	if err != nil {
		return err
	}
	last := len(DB.Products) - 1
	DB.Products[idx] = DB.Products[last]
	DB.Products = DB.Products[:last]
	return nil
}

func findProductByID(id int) (int, error) {
	for idx, p := range DB.Products {
		if id == p.ID {
			return idx, nil
		}
	}
	return 0, constant.ErrProductNotFound
}
