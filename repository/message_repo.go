package repository

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/models"
)

func SaveMessage(msg *models.EmpxMessage) error {
	// 注意 postgres sql 插入时，表要标明在哪个域（public）内
	query := `INSERT INTO public.packets (node_id, type, message, received_at) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, msg.NodeID, msg.Type, msg.Value, msg.TS)
	return err
}
