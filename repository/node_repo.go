package repository

import (
	"EmqxBackEnd/database"
	"log"
)

// 根据节点id获取用户id
func GetUserIdByNodeId(nodeId int) (int, error) {
	query := `select user_id from public.node where id=$1`
	var userId int
	err := database.DB.QueryRow(query, nodeId).Scan(&userId)
	if err != nil {
		log.Println("GetUserIdByNodeId error:", err)
		return 0, err
	}
	return userId, nil
}
