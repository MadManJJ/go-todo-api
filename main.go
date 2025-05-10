package main

import (
	"net/http"
	"os"
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

	// * Todo
	router.GET("/todos", getTodosHandler)
	router.GET("/users/:userId/todos", getTodosHandler) // * get todo by userId
	router.GET("/todos/:id", getTodoHandler)
	router.POST("/todos", createTodoHandler)
	router.PUT("/todos/:id", updateTodoHandler)
	router.DELETE("/todos/:id", deleteTodHandler)

	// * Auth
	router.POST("/auth/register", createUserHandler)

	host := os.Getenv("HOST")
	if host == "" {
			host = "localhost" // default to all interfaces
	}

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080" // default to 8080 if not set
	}

	router.Run(host + ":" + port)
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
	pageStr := c.DefaultQuery("page", "0")
	limitStr := c.DefaultQuery("limit", "0")

	page, err := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)

	if err != nil || err2 != nil || page < 0 || limit < 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   "Invalid pagination parameters",
		})
		return
	}
	offset := (page - 1) * limit

	userId := c.Param("userId") // * optional
	var todos []models.Todo

	if userId != "" {
		userIdUint64, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "failed",
				"error":   "Invalid user ID",
			})
			return
		}
		todos = todo.GetTodos(gormdb, limit, offset, uint(userIdUint64))
	} else {
		todos = todo.GetTodos(gormdb, limit, offset)
	}

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

func updateTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   "Invalid ID",
		})
		return
	}

	updatedTodo := new(models.Todo)

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed",
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	updatedTodo.ID = uint(id)

	if updatedTodo, err = todo.UpdateTodo(gormdb, updatedTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed",
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "success",
		"data" : updatedTodo,
	})
}

func deleteTodHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "failed",
			"error" : "Invalid ID",
		})
		return
	}

	if err = todo.DeleteTodo(gormdb, uint(id)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "failed",
			"error" : err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "success",
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
