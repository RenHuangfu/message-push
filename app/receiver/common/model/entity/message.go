package entity

import "time"

type SendMessageParam struct {
	AppID      int32
	ClientIDs  []int32
	Content    string
	Delay      int64
	SendTime   time.Time
	SendCount  int32
	Status     int32
	IsDel      int32
	CreateTime string
	UpdateTime string
}

type SendMessageResult struct {
	MessageID int32
}
