package db

import (
	"log"

	"github.com/nicopv-dev/go-first-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=1999_nico dbname=go_first_api_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}else {
		log.Println("DB Connection Success...")
	}

	return db
}

func Migrate() {
	db := DBConnection()

	log.Println("Migrating...")

	db.AutoMigrate(&models.User{})
}