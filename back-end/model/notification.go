package model

import "time"

type Notification struct {
	NotifyId string    `json:"notifyId"`
	UserId   string    `json:"userId"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Status   string    `json:"status"`
	ReadAt   time.Time `json:"readAt"`
}
