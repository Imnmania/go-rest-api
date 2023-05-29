package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables!!!")
		os.Exit(2)
	}
	log.Println("Environment variables loaded...")
}
