package entity

import (
	"github.com/gorilla/websocket"
	"sync"
)

type ClientRequest struct {
	AppId    string `json:"app_id"`
	ClientId string `json:"client_id"`
}

type ClientKey struct {
	AppId    string
	ClientId string
}

type Message struct {
	AppID     int    `json:"app_id"`
	ClientIDs []int  `json:"client_ids"`
	Content   string `json:"content"`
	//SendTime  time.Time `json:"send_time"`
}

type Connect struct {
	sync.Mutex
	Conn *websocket.Conn
}
