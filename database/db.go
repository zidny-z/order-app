package database

import "gorm.io/gorm"

type Database struct {
	db *gorm.DB
}