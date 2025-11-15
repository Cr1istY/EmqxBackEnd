package repository

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/models"
	"database/sql"
	"log"
)

// SaveMessage 保存消息
func SaveMessage(msg *models.EmpxMessage) error {
	// 注意 postgres sql 插入时，表要标明在哪个域（public）内
	query := `INSERT INTO public.packets (node_id, type, message, received_at) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, msg.NodeID, msg.Type, msg.Value, msg.TS)
	return err
}

// GetMessages 读取数据库
func GetMessages(msgType string) ([]models.EmpxMessage, error) {
	query := `SELECT node_id, type, message, received_at FROM public.packets WHERE type = $1`
	rows, err := database.DB.Query(query, msgType)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("rows.Close() failed: %v", err)
		}
	}(rows)
	var messages []models.EmpxMessage
	for rows.Next() {
		var msg models.EmpxMessage
		err := rows.Scan(&msg.NodeID, &msg.Type, &msg.Value, &msg.TS)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
