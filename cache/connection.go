package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Rdb *redis.Client
	Ctx context.Context
}

var cache Cache

func Init() (err error) {

	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		return err
	}

	cache.Rdb = redis.NewClient(opt)
	cache.Ctx = context.Background()
	err = cache.Rdb.Ping(cache.Ctx).Err()

	return err
}

func GetCache() Cache {
	return cache
}
