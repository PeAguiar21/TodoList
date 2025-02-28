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
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	database.ConnectMongoDB(mongoURI)

	collection := database.GetCollection(dbName, "tasks")
	log.Printf("Coleção %s obtida com sucesso", collection.Name())
}
