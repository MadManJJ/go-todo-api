package auth

import (
	"github.com/MadManJJ/go-todo-api/models"
	"gorm.io/gorm"

	// "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}