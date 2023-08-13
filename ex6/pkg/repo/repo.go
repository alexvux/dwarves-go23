package repo

import (
	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
)

type Database struct {
	Products []model.Product
	Cart     []model.Item
}

var LocalDB *Database

func init() {
	LocalDB = newDatabase()
}

func newDatabase() *Database {
	return &Database{
		Products: []model.Product{},
		Cart:     []model.Item{},
	}
}
