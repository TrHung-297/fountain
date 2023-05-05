

package main

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/grand"
	"github.com/TrHung-297/fountain/baselib/redis_client"
)

func main() {
	configMap := make(map[string]interface{})
	myConfig := redis_client.RedisConfig{
		Name:         "cache",
		Addrs:        []string{"127.0.0.1:6379"},
		Idle:         100,
		Active:       100,
		DialTimeout:  1,
		ReadTimeout:  1,
		WriteTimeout: 1,
		IdleTimeout:  10,

		DBNum:    0,
		Password: "bitnami",
	}

	configMap["Redis"] = []redis_client.RedisConfig{myConfig}
	// viper.Set("config.json", base.JSONDebugDataString(configMap))

	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	env := viper.GetString("Env.Environment")
	log.Printf("Env: %s", env)

	envMap := viper.GetStringMap("Env")
	log.Printf("envMap: %+v", base.JSONDebugDataString(envMap))

	redisMap := viper.GetStringMap("Redis")
	log.Printf("redisMap: %+v", base.JSONDebugDataString(redisMap))

	redis_client.InstallRedisClientManager()

	myPo := redis_client.GetRedisClient("cache")

	myRand := grand.RandomAlphaOrNumeric(100, true, true)

	res, err := myPo.Get().SAdd(context.Background(), "test:1", myRand, 0).Result()
	g_log.V(3).Infof("Set res: %+v, err: %+v", res, err)

	resPop, errPop := myPo.Get().SPop(context.Background(), "test:1").Result()
	g_log.V(3).Infof("Get resPop: %+v, errPop: %+v", resPop, errPop)

}
