package redis_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/TrHung-297/fountain/baselib/cache"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/redis_client"
)

var (
	defaultTimeout = 5 * time.Second
)

var (
	// DefaultKey the collection name of redis for cache adapter.
	DefaultKey = "gtvCacheRedis"
)

// Cache is Redis cache adapter.
type Cache struct {
	redisPool *redis_client.RedisPool // redis connection pool

	name        string
	key         string
	configRedis []*redis_client.RedisConfig
}

// NewRedisCache create new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &Cache{key: DefaultKey}
}

// associate with config key.
func (rc *Cache) associate(originKey interface{}) string {
	return fmt.Sprintf("%s:%s", rc.key, originKey)
}

// Get cache from redis.
func (rc *Cache) Get(key string) interface{} {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var err error
	if v, err := conn.Get(ctx, key).Result(); err == nil {
		return v
	}

	g_log.V(3).Infof("Cache get err: %+v", err)
	return nil
}

// GetMulti get cache from redis.
func (rc *Cache) GetMulti(keys []string) []interface{} {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	var args []string
	for _, key := range keys {
		args = append(args, rc.associate(key))
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	values, err := conn.MGet(ctx, args...).Result()
	if err != nil {
		return nil
	}

	g_log.V(3).Infof("Cache GetMulti err: %+v", err)
	return values
}

// Put put cache to redis.
func (rc *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := conn.Set(ctx, key, val, timeout).Err()

	if err != nil {
		g_log.V(3).Infof("Cache Put err: %+v", err)
	}
	return err
}

// PutWithoutExprise func;
// Put without exprise
func (rc *Cache) PutWithoutExprise(key string, val interface{}, timeout time.Duration) error {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := conn.Set(ctx, key, val, 0).Err()
	if err != nil {
		g_log.V(3).Infof("Cache PutWithoutExprise err: %+v", err)
		return err
	}

	return err
}

// Delete delete cache in redis.
func (rc *Cache) Delete(key string) error {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}
	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := conn.Del(ctx, key).Err()

	if err != nil {
		g_log.V(3).Infof("Cache Delete err: %+v", err)
	}
	return err
}

// IsExist check cache's existence in redis.
func (rc *Cache) IsExist(key string) bool {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	v, err := conn.Exists(ctx, key).Result()
	if err != nil {
		g_log.V(3).Infof("Cache IsExist err: %+v", err)
		return false
	}
	return v > 0
}

// Incr increase counter in redis.
func (rc *Cache) Incr(key string) error {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := conn.Incr(ctx, key).Err()
	if err != nil {
		g_log.V(3).Infof("Cache Incr err: %+v", err)
	}
	return err
}

// Decr decrease counter in redis.
func (rc *Cache) Decr(key string) error {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := conn.IncrBy(ctx, key, -1).Err()
	if err != nil {
		g_log.V(3).Infof("Cache Decr err: %+v", err)
	}
	return err
}

// ClearAll clean all cache in redis. delete this redis collection.
func (rc *Cache) ClearAll() error {
	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}

	conn := rc.redisPool.Get()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	cachedKeys, err := conn.Keys(ctx, rc.key+":*").Result()
	if err != nil {
		return err
	}

	for _, str := range cachedKeys {
		ctxChild, cancelChild := context.WithTimeout(context.Background(), defaultTimeout)
		if err = conn.Del(ctxChild, str).Err(); err != nil {
			g_log.V(3).Infof("Cache Decr err: %+v", err)
			cancelChild()
			return err
		}

		cancelChild()
	}

	return err
}
