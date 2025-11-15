package repository

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/models"
	"database/sql"
	"errors"
	"log"
)

func GetAdminByUser(username string) (*models.EmpxAdmin, error) {
	query := `SELECT id, username, password, status, created_time FROM public.admin WHERE username = $1`
	var admin models.EmpxAdmin
	err := database.DB.QueryRow(query, username).Scan(
		&admin.ID,
		&admin.Username,
		&admin.Password,
		&admin.Status,
		&admin.CreatedTime,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("No admin found with username:", username)
			return nil, nil
		}
		log.Println("Error querying admin:", err)
		return nil, err
	}
	return &admin, nil

}
