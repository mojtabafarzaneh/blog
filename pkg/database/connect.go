package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dsn := "host=localhost user=postgres password=mojtaba7878 dbname=blog port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("couldn't connect to the database", err)
	}

	var DB *gorm.DB

	DB = db

	return DB
}
