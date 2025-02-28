package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbName    string
	DebugMode bool
	MongoURI  string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error to load the file .env: %v", err)
	}

	DbName = os.Getenv("DB_NAME")
	MongoURI = os.Getenv("MONGO_URI")
}
