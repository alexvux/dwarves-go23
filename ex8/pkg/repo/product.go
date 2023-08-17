package repo

import (
	"github.com/alexvux/dwarves-go23/ex8/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex8/pkg/model"
)

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := DB.Find(&products).Error
	if err == nil && len(products) == 0 {
		return products, constant.ErrEmptyProductList
	}
	return products, err
}

func AddProduct(product model.Product) error {
	if _, err := getProductByID(int(product.ID)); err == nil {
		return constant.ErrProductAlreadyExist
	}
	return DB.Create(&product).Error
}

func UpdateProduct(product model.Product) error {
	id := int(product.ID)
	if _, err := getProductByID(id); err != nil {
		return constant.ErrProductNotFound
	}
	return DB.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
}

func DeleteProduct(id int) error {
	if _, err := getProductByID(id); err != nil {
		return constant.ErrProductNotFound
	}
	return DB.Delete(&model.Product{}, id).Error
}

func getProductByID(id int) (model.Product, error) {
	var product model.Product
	err := DB.First(&product, id).Error
	return product, err
}
