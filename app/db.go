package app

import (
	"fmt"

	"github.com/mkamadeus/iot-smart-retail/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
