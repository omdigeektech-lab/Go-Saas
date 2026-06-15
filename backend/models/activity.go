package models

import "time"

type Activity struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"taskId"`
	Action    string    `json:"action"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
