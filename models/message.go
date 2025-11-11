package models

type EmpxMessage struct {
	NodeID int         `json:"nodeID"`
	Type   string      `json:"type"`
	Value  interface{} `json:"value"`
	TS     int64       `json:"ts"`
}
