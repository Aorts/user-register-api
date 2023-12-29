package infrastructures

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedis(url, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic("error connect redis client: " + err.Error())
	}

	return client
}
