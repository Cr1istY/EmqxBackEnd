package service

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"
)

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
