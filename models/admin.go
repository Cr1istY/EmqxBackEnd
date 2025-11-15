package models

import (
	"time"
)

type EmpxAdmin struct {
	ID          int       `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password"`
	Status      int8      `json:"status" db:"status"`
	CreatedTime time.Time `json:"createdTime" db:"created_time"`
}
