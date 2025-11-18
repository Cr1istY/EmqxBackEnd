package service

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"
)

func SaveNode(node *models.Node) error {
	return repository.SaveNode(node.ID, node.UserId)
}
