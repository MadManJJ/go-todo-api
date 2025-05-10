package main

import (
	"net/http"
	"strconv"

	"github.com/MadManJJ/go-todo-api/auth"
	"github.com/MadManJJ/go-todo-api/db"
	"github.com/MadManJJ/go-todo-api/models"
	"github.com/MadManJJ/go-todo-api/todo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func main() {
	gormdb = db.InitDB()
	router := gin.Default()

	router.GET("/todos", getTodosHandler)
	router.GET("/todos/:id", getTodoHandler)
	router.POST("/todos", createTodoHandler)
	router.POST("/auth/register", createUserHandler)

	router.Run("localhost:8080")
}

func createTodoHandler(c *gin.Context) {
	newTodo := new(models.Todo)

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed",
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	if err := todo.CreateTodo(gormdb, newTodo); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed to create todo",
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "success",
	})
}

func getTodosHandler(c *gin.Context) {
	todos := todo.GetTodos(gormdb)
	if(len(todos) == 0){
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "failed",
			"error": "No todos found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{ 
		"message" : "success",
		"count" : len(todos),
		"data" : todos,
	})
}

func getTodoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   "Invalid todo ID: " + err.Error(),
		})
		return
	}

	todo, err := todo.GetTodo(gormdb, uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "failed",
			"error": "No todos found: " + err.Error(),
		})
		return
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{ 
		"message" : "success",
		"data" : todo,
	})
}

func createUserHandler(c *gin.Context) {
	newUser := new(models.User)

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed",
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	if err := auth.CreateUser(gormdb, newUser); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed to create todo",
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "success",
	})
}