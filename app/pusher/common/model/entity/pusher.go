package entity

type PusherJson struct {
	Host        string  `json:"host"`
	Port        int     `json:"port"`
	Topic       string  `json:"topic"`
	Weight      int     `json:"weight"`
	Connections int     `json:"connections"`
	CPULoad     float64 `json:"cpu_load"`
	UpdatedAt   int64   `json:"updated_at"`
}
