package models

type EmpxMessage struct {
	NodeID int         `json:"nodeID" db:"node_id"`
	Type   string      `json:"type" db:"type"`
	Value  interface{} `json:"value" db:"message"`
	TS     int64       `json:"ts" db:"received_at"`
}
