package todo

import (
	"log"

	"github.com/MadManJJ/go-todo-api/models"
	"gorm.io/gorm"
)

func CreateTodo(db *gorm.DB, todo *models.Todo) error {
	result := db.Create(todo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetTodos(db *gorm.DB) []models.Todo {
	var todos []models.Todo
	result := db.Find(&todos)

	if result.Error != nil {
		log.Fatalf("Error get books: %v", result.Error)
	}

	return todos
}