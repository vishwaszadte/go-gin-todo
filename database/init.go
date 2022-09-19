package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "postgres://hcglesuf:NIgWVpgjOCT2uM70yXGO6okT97Wbc2kk@jelani.db.elephantsql.com/hcglesuf"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d;
}

func GetDB() *gorm.DB {
	return db
}