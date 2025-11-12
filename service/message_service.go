package service

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"
)

func ProcessEmpxMessage(msg *models.EmpxMessage) error {
	return repository.SaveMessage(msg)
}
