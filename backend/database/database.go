package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init opens a SQLite database at the provided path and runs automigrations
// for the supplied models.
func Init(path string, models ...interface{}) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := database.AutoMigrate(models...); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	db = database
	return db, nil
}

// DB returns the active gorm database handle.
func DB() *gorm.DB {
	return db
}
