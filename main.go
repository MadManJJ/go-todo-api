package main

import (
	"net/http"

	"github.com/MadManJJ/go-todo-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "HI")
	})

	router.Run("localhost:8080")
}