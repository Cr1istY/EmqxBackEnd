package models

type Node struct {
	ID     int `json:"id" db:"id"`
	UserId int `json:"userId" db:"user_id"`
}
