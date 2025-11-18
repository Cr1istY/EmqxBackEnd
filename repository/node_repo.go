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

// SaveNode 绑定节点
func SaveNode(nodeId, userId int) error {
	query := `insert into public.node(id, user_id) values($1, $2)`
	_, err := database.DB.Exec(query, nodeId, userId)
	return err
}

func UpdateNode(nodeId, userId int) error {
	query := `update public.node set user_id=$1 where id=$2`
	_, err := database.DB.Exec(query, userId, nodeId)
	return err
}

func CheckNode(nodeId int) (bool, error) {
	query := `select exists(select 1 from public.node where id=$1)`
	var exists bool
	err := database.DB.QueryRow(query, nodeId).Scan(&exists)
	if err != nil {
		log.Println("CheckNode error:", err)
		return false, err
	}
	return exists, nil
}
