package external

import (
	"fmt"

	"github.com/mkamadeus/iot-smart-retail/config"
	"github.com/mkamadeus/iot-smart-retail/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{}, &models.Item{}, &models.Transaction{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
