package main

import (
	"net/http"

	"github.com/MadManJJ/go-todo-api/db"
	"github.com/MadManJJ/go-todo-api/todo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func main() {
	gormdb = db.InitDB()
	router := gin.Default()

	router.GET("/todos", getTodosHandler)

	router.Run("localhost:8080")
}

func getTodosHandler(c *gin.Context) {
	todos := todo.GetTodos(gormdb)
	if(len(todos) == 0){
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No todos found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}