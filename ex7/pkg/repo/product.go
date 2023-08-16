package repo

import (
	"github.com/alexvux/dwarves-go23/ex7/pkg/constant"
	"github.com/alexvux/dwarves-go23/ex7/pkg/model"
)

func GetAllProducts() ([]model.Product, error) { // what happen if no product in db?
	var products []model.Product
	err := DB.Find(&products).Error
	return products, err
}

func GetProductByID(id int) (model.Product, error) { // what happen if no product in db?
	var product model.Product
	err := DB.First(&product, id).Error
	return product, err
}

func AddProduct(product model.Product) error { // do 3 fields created_at, updated_at, deleted_at are created automatically?
	var p model.Product
	if err := DB.First(&p, product.ID).Error; err == nil {
		return constant.ErrProductAlreadyExist
	}
	return DB.Create(&product).Error
}

func UpdateProduct(id int, product model.Product) error { // what happen if no product in db?
	var p model.Product
	// err := DB.First(&p, id).Error
	// if err != nil {
	// 	return err
	// }
	return DB.Model(&p).Updates(product).Error
}

func DeleteProduct(id int) error { // what happen if no product in db?
	return DB.Delete(&model.Product{}, id).Error
}
