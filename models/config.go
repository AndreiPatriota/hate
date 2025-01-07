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

type Poke struct {
	gorm.Model
	Nome string
	Tipo string
	Caracteristica string
	Peso float32
	FotoUrl string
	ChoroUrl string
}

var DB *gorm.DB
func InitDb() {
	db, err := gorm.Open(sqlite.Open("./models/base.db"), &gorm.Config{})
	if err != nil {
		panic("Probelma na conexão com o Banco!")
	}

	db.AutoMigrate(&Produto{})
	db.AutoMigrate(&Poke{})

	DB = db
}