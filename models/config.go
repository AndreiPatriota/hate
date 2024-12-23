package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome string
	Descricao string
}

var DB *gorm.DB
func InitDb() {
	db, err := gorm.Open(sqlite.Open("./models/base.db"), &gorm.Config{})
	if err != nil {
		panic("Probelma na conexão com o Banco!")
	}

	db.AutoMigrate(&Produto{})

	DB = db
}