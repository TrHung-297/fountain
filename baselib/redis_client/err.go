
package redis_client

import (
	redis "github.com/go-redis/redis/v8"
)

var (
	ErrNil    = redis.Nil
	ErrClosed = redis.ErrClosed
)
