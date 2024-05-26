package database

import (
	"crud-restapi/models"
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_CONNECTION")), &gorm.Config{})
	if err != nil {
		return nil, errors.New("database connection failed")
	}

	db.AutoMigrate(&models.Task{}, &models.User{})

	return db, nil
}
