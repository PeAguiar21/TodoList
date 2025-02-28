package routes

import (
	"github.com/PeAguiar21/TodoList/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/task", handlers.CreateTask)
}
