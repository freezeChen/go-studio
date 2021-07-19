package db

type Config struct {
	DriverName string `json:"driver_name"`
	Source     string `json:"source"`
	Show       bool   `json:"show"`
	Max        int    `json:"max"`
	Idle       int    `json:"idle"`
}
