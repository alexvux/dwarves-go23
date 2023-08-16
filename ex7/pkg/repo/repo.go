package repo

import (
	"github.com/alexvux/dwarves-go23/ex7/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(connStr string) error {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
