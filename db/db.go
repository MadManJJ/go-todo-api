package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/MadManJJ/go-todo-api/models"
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT")) // * change string to int
	user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASSWORD")
  dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

		  // New logger for detailed SQL logging
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold: time.Second, // Slow SQL threshold
      LogLevel:      logger.Info, // Log level
      Colorful:      true,        // Enable color
    },
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: newLogger,
  })

  if err != nil {
    panic("failed to connect to database")
  }
	fmt.Println(db)

  // Migrate the schema
  db.AutoMigrate(&models.User{}, &models.Todo{})
}