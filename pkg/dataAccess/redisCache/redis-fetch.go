package redisCache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"

	"project-tenant/pkg/dataAccess"
)

type redisF struct {
	cacheCollection *redis.Client
}

func NewredisF() *redisF {
	return &redisF{
		cacheCollection: dataAccess.Cache,
	}
}

func (cache *redisF) Get(ctx context.Context, key string) *redis.StringCmd {

	return dataAccess.Cache.Get(ctx, key)

}
func (cache *redisF) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return dataAccess.Cache.Set(ctx, key, value, expiration).Err()
}
