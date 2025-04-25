package entity

import "time"

type ClientRequest struct {
	AppId    string `json:"app_id"`
	ClientId string `json:"client_id"`
	Type     string `json:"type"`
}

type ClientKey struct {
	AppId    string
	ClientId string
	Type     string
}

type Message struct {
	AppID    int       `json:"app_id"`
	ClientID int       `json:"client_id"`
	Content  string    `json:"content"`
	SendTime time.Time `json:"send_time"`
}
