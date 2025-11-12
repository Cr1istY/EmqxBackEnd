package repository

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/models"
)

// 在前端处理
//type MessageType string
//
//const (
//	Register    MessageType = "0"
//	Confirm     MessageType = "1"
//	Temperature MessageType = "2"
//	Humidity    MessageType = "3"
//	MQ6         MessageType = "4"
//)
//
//var messageTypeMap = map[MessageType]string{
//	Register:    "注册",
//	Confirm:     "确认",
//	Temperature: "温度",
//	Humidity:    "湿度",
//	MQ6:         "天然气",
//}

func SaveMessage(msg *models.EmpxMessage) error {
	// 注意 postgres sql 插入时，表要标明在哪个域（public）内
	query := `INSERT INTO public.packets (node_id, type, message, received_at) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, msg.NodeID, msg.Type, msg.Value, msg.TS)
	return err
}

// 读取数据库
func GetMessages(msgType string) ([]models.EmpxMessage, error) {
	query := `SELECT node_id, type, message, received_at FROM public.packets WHERE type = $1`
	rows, err := database.DB.Query(query, msgType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
