package redisCache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type CacheMethods interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
