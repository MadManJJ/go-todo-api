package todo

import (
	"log"

	"github.com/MadManJJ/go-todo-api/models"
	"gorm.io/gorm"
)

func GetTodos(db *gorm.DB) []models.Todo {
	var todos []models.Todo
	result := db.Find(&todos)

	if result.Error != nil {
		log.Fatalf("Error get books: %v", result.Error)
	}

	return todos
}