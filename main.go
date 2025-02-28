package main

import (
	"log"

	"github.com/PeAguiar21/TodoList/config"
	"github.com/PeAguiar21/TodoList/database"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load file .env")
	}

	err = database.ConnectMongoDB(config.MongoURI)
	if err != nil {
		log.Fatal("Error to connect to mongoDB: ", err.Error())
	}

}
