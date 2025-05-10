package entity

import "time"

type SearchPusherParam struct {
	AppId    string
	ClientId string
}

type Pusher struct {
	Host        string
	Port        int
	Topic       string
	Weight      int       // 初始权重
	Connections int       // 当前连接数
	CPULoad     float64   // CPU负载
	LastUpdate  time.Time // 上次更新时间
}

type PusherJson struct {
	Host        string  `json:"host"`
	Port        int     `json:"port"`
	Topic       string  `json:"topic"`
	Weight      int     `json:"weight"`
	Connections int     `json:"connections"`
	CPULoad     float64 `json:"cpu_load"`
	UpdatedAt   int64   `json:"updated_at"`
}
