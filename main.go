package main

import (
	"log"
	"os"

	"github.com/PeAguiar21/TodoList/database"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load file .env")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	err = database.ConnectMongoDB(mongoURI)
	if err != nil {
		log.Fatal("Error to connect to mongoDB: ", err.Error())
	}

	collection := database.GetCollection(dbName, "tasks")
	log.Printf("Collection %s get with sucess", collection.Name())
}
