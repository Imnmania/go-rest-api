package config

import (
	"log"
	"os"

	"github.com/imnmania/go_fiber_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database!!!")
		os.Exit(2)
	}

	log.Println("Database connection established...")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations...")
	runMigrations()
}

func runMigrations() {
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)
}
