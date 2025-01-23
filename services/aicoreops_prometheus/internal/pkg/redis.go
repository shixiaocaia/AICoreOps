package pkg

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis(c string) redis.Cmdable {
	client := redis.NewClient(&redis.Options{
		Addr: c,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}
