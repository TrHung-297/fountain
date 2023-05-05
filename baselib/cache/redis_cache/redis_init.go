package redis_cache

import (
	"encoding/json"
	"fmt"

	"github.com/TrHung-297/fountain/baselib/cache"
	"github.com/TrHung-297/fountain/baselib/redis_client"
)

// StartAndGC start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *Cache) StartAndGC(configsJSON ...string) error {
	configFromJSON := make([]*redis_client.RedisConfig, 0)

	for _, configJSON := range configsJSON {
		conf := &redis_client.RedisConfig{}

		err := json.Unmarshal([]byte(configJSON), conf)
		if configJSON == "" || err != nil {
			continue
		}

		configFromJSON = append(configFromJSON, conf)
	}

	if len(configFromJSON) == 0 {
		return fmt.Errorf("not found any config")
	}

	rc.configRedis = configFromJSON
	for _, config := range configFromJSON {
		rc.name = config.Name
		if rc.name != "" {
			break
		}
	}

	rc.connectInit()

	return nil
}

// connect to redis.
func (rc *Cache) connectInit() {
	redis_client.InstallRedisClientManagerWithConfig(rc.configRedis)
	rc.redisPool = redis_client.GetRedisClient(rc.name)

	if rc.redisPool == nil {
		err := fmt.Errorf("Can't connect to cache %s", rc.name)
		panic(err)
	}
}

func init() {
	cache.Register("redis", NewRedisCache)
}
