package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func GetConfigDB() Config {
	GetConfig()
	return Config{
		DB_USERNAME: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
}

func GetConfigPort() string {
	GetConfig()
	return ":" + os.Getenv("WEBSITE_PORT")
}
