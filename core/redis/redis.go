package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	client *redis.Client
}

func InitRedis(ctx context.Context, c Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Username: c.UserName,
		Password: c.Password,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return &Redis{client: client}
}

func (r *Redis) GetConn(ctx context.Context) *redis.Conn {
	return r.client.Conn(ctx)
}

func (r *Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Del(ctx context.Context, key ...string) error {
	return r.client.Del(ctx, key...).Err()
}
