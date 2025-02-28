package main

import (
	"log"

	"github.com/PeAguiar21/TodoList/api/routes"
	"github.com/PeAguiar21/TodoList/config"
	"github.com/PeAguiar21/TodoList/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	err := database.ConnectMongoDB(config.MongoURI)
	if err != nil {
		log.Fatalln("Error to connect to mongoDB: ", err.Error())
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Erro ao iniciar o servidor:", err)
	}

}
