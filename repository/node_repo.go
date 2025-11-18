package repository

import (
	"EmqxBackEnd/database"
	"log"
)

// GetUserIdByNodeId 根据节点id获取用户id
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

// 绑定节点
func SaveNode(nodeId, userId int) error {
	query := `insert into public.node(id, user_id) values($1, $2)`
	_, err := database.DB.Exec(query, nodeId, userId)
	return err
}
