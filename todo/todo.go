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

func GetTodos(db *gorm.DB, limit int, offset int, userId ...uint) []models.Todo {
	var todos []models.Todo
	var result *gorm.DB
	
	query := db
	if offset != 0 {
		query = db.Limit(limit).Offset(offset)
	}
	if len(userId) > 0 {
		result = query.Where("user_id = ?", userId[0]).Find(&todos)
	}else{
		result = query.Find(&todos)
	}

	if result.Error != nil {
		log.Fatalf("Error get books: %v", result.Error)
	}

	return todos
}

func GetTodo(db *gorm.DB, id uint) (*models.Todo, error) {
	var todo models.Todo
	result := db.First(&todo, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}