package repository

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/models"
)

type MessageType string

const (
	Register    MessageType = "0"
	Confirm     MessageType = "1"
	Temperature MessageType = "2"
)

var messageTypeMap = map[MessageType]string{
	Register:    "注册",
	Confirm:     "确认",
	Temperature: "温度",
}

func SaveMessage(msg *models.EmpxMessage) error {
	// 注意 postgres sql 插入时，表要标明在哪个域（public）内
	query := `INSERT INTO public.packets (node_id, type, message, received_at) VALUES ($1, $2, $3, $4)`
	realType := "未知"
	if desc, exists := messageTypeMap[MessageType(msg.Type)]; exists {
		realType = desc
	}
	_, err := database.DB.Exec(query, msg.NodeID, realType, msg.Value, msg.TS)
	return err
}
