package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tasks struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TaskId      int                `json:"taskId" bson:"taskId"`
	Name        string             `json:"name" bson:"name"`
	CreatedDate time.Time          `json:"createdDate" bson:"createdDate"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updatedDate"`
	ToDoDate    time.Time          `json:"toDoDate" bson:"toDoDate"`
	IsConclued  bool               `json:"isConclued" bson:"isConclued"`
}
