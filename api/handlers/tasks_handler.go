package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/PeAguiar21/TodoList/config"
	"github.com/PeAguiar21/TodoList/database"
	"github.com/PeAguiar21/TodoList/models"
	"github.com/gin-gonic/gin"
)

// func GetTask() {

// }

func CreateTask(c *gin.Context) {
	var task models.Tasks

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection(config.DbName, "tasks")

	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Println("Error to include task:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to include task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          result.InsertedID,
		"taskId":      task.TaskId,
		"name":        task.Name,
		"createdDate": task.CreatedDate,
		"updatedDate": task.UpdatedDate,
		"toDoDate":    task.ToDoDate,
		"isConclued":  task.IsConclued,
	})

}

// func UpdateTask() {

// }

// func DeleteTask() {

// }
