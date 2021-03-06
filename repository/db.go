package repository

import (
	"fmt"

	"github.com/mkamadeus/iot-smart-retail/config"
	"github.com/mkamadeus/iot-smart-retail/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// setup join table
	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Transaction{},
		&models.Order{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
