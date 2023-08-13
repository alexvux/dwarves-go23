package repo

import (
	"github.com/alexvux/dwarves-go23/ex6/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
)

func GetAllProducts(prodcut model.Product) []model.Product {
	return LocalDB.Products
}

func AddProduct(product model.Product) error {
	if _, err := findProductByID(product.ID); err != nil {
		return err
	}
	LocalDB.Products = append(LocalDB.Products, product)
	return nil
}

func UpdateProduct(product model.Product) error {
	idx, err := findProductByID(product.ID)
	if err != nil {
		return err
	}
	LocalDB.Products[idx] = product
	return nil
}

func DeleteProduct(product model.Product) error {
	idx, err := findProductByID(product.ID)
	if err != nil {
		return err
	}
	last := len(LocalDB.Products) - 1
	LocalDB.Products[idx] = LocalDB.Products[last]
	LocalDB.Products = LocalDB.Products[:last]
	return nil
}

func findProductByID(id int) (int, error) {
	for i, p := range LocalDB.Products {
		if id == p.ID {
			return i, nil
		}
	}
	return 0, constant.ErrProductNotFound
}
