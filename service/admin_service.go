package service

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"

	"github.com/golang-jwt/jwt/v5"

	"time"
)

var jwtSecret = []byte("cqupt") // Should be stored securely

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func CheckLogin(username, password string) bool {
	var user *models.EmpxAdmin
	user, err := repository.GetAdminByUser(username)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	if user.Status != 1 {
		return false
	}
	if user.Password != password {
		return false
	}
	return true
}
