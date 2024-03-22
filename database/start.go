package database

import (
	"fmt"
	"order-app/config"
	"order-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() (Database, error) {
	dbConfig := config.GetConfigDB()

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DB_HOST, dbConfig.DB_USERNAME, dbConfig.DB_PASSWORD, dbConfig.DB_NAME, dbConfig.DB_PORT)

	db, error := gorm.Open((postgres.Open(connection)), &gorm.Config{})
	if error != nil {
		fmt.Println("Error connecting to database")
		return Database{}, error
	}

	// migrate
	if error := db.Debug().AutoMigrate(&models.Order{}, &models.Item{}); error != nil {
		fmt.Println("Error migrating database")
		return Database{}, error
	}

	fmt.Println("Database connected")
	return Database{db: db}, nil
}