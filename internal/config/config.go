package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var APIHost string
var ServiceHost string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("File download error .env: %v", err)
	}
	APIHost = os.Getenv("API_HOST")
	ServiceHost = os.Getenv("SERVICE_HOST")
}
