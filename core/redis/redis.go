package redis

import "github.com/go-redis/redis/v8"

type Redis struct {
}

func InitRedis(c Config) *Redis {
	redis.NewClusterClient()
	redis.NewSentinelClient()
}
