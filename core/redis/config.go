package redis

type Config struct {
	Address  string `json:"address"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Db       string `json:"db"`
}
