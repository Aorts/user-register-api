package infrastructures

import (
	"playground/grpc/config"

	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Host,
	})
}
