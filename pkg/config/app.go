package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost port=5432 user=postgres dbname=go_books password=123qwe sslmode=disable"

	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err) 
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}