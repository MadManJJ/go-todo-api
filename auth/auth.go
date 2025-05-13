package auth

import (
	"os"
	"time"

	"github.com/MadManJJ/go-todo-api/models"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v4"
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

func LoginUser(db *gorm.DB, user *models.User) (string, error) {
	var selectedUser models.User
	result := db.Where("email = ?", user.Email).First(&selectedUser)
	if result.Error != nil {
		return "", result.Error // * Return error if no user with the provided email is found
	}

	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
  claims["user_id"] = selectedUser.ID
  claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
  if err != nil {
    return "", err
  }

	return t, nil
}