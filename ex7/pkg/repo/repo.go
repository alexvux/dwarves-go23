package repo

import (
	"log"

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
	log.Println("Connect to postgres DB successfully")

	err = db.AutoMigrate(&model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		return err
	}
	log.Println("Auto migrate successfully")

	DB = db
	return nil
}
