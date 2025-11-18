package models

import "time"

type EmpxMessage struct {
	NodeID int       `json:"nodeID" db:"node_id"`
	Type   string    `json:"type" db:"type"`
	Value  string    `json:"value" db:"message"`
	TS     time.Time `json:"ts" db:"received_at"`
	UserId int       `json:"userId" db:"user_id"`
}

type GetMessage struct {
	UserId int `json:"userId" db:"user_id"`
	Type   int `json:"type" db:"type"`
}
