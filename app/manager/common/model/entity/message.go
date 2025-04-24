package entity

import "time"

type Message struct {
	AppID    int       `json:"app_id"`
	ClientID int       `json:"client_id"`
	Content  string    `json:"content"`
	SendTime time.Time `json:"send_time"`
}
