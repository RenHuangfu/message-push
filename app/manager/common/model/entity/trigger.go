package entity

import "time"

type TriggerParam struct {
	AppID     int32
	SendTime  time.Time
	SendCount int32
	Delay     time.Duration
}
